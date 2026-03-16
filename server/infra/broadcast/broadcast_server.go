package broadcast

import (
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/codec"
	"fmt"
	"github.com/charmbracelet/log"
	"net"
	"sync"
)

type ServerConfig struct {
	port int
}

func NewBroadcastServerConfig() ServerConfig {
	return ServerConfig{
		port: 8666,
	}
}

type Handlers map[socket.PayloadType][]func(broadcast socket.Message) *socket.Message

type Server struct {
	listening         bool
	conn              *net.UDPConn
	broadcastHandlers Handlers
	codec             *codec.CborCodec
	port              int

	clientsMux  sync.RWMutex
	clients     []*net.UDPAddr
	clientsChan chan *net.UDPAddr
}

func NewServer(config ServerConfig) *Server {
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", config.port))
	if err != nil {
		log.Fatal("Failed to resolve server address:", err)
	}
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatal("Failed to start UDP server:", err)
	}

	return &Server{
		conn:              conn,
		codec:             codec.NewCborCodec(),
		broadcastHandlers: make(Handlers),
		port:              config.port,

		clients:     make([]*net.UDPAddr, 0),
		clientsChan: make(chan *net.UDPAddr),
	}
}

func (s *Server) OnBroadcast(
	payloadType socket.PayloadType,
	handleFunc func(socket.Message) *socket.Message,
) {
	s.broadcastHandlers[payloadType] = append(s.broadcastHandlers[payloadType], handleFunc)
}

func (s *Server) Start() {
	log.Infof("Starting Broadcast server on port: %d", s.port)
	s.listening = true

	go s.listen()
	go s.addClients()
}

func (s *Server) listen() {
	buffer := make([]byte, 1024)
	for {
		if !s.listening {
			return
		}

		n, clientAddr, err := s.conn.ReadFromUDP(buffer)
		s.clientsChan <- clientAddr

		if err != nil {
			log.Error("Failed to read from UDP:", err)
			continue
		}
		go s.handlePacket(buffer[:n], clientAddr)
	}
}

func (s *Server) addClients() {
outer:
	for {
		select {
		case newClientAddress := <-s.clientsChan:
			s.clientsMux.Lock()
			for _, client := range s.clients {
				if client.String() == newClientAddress.String() {
					s.clientsMux.Unlock()
					continue outer
				}
			}
			s.clients = append(s.clients, newClientAddress)
			s.clientsMux.Unlock()
		}
	}
}

func (s *Server) Stop() {
	log.Info("Stopping Broadcast server...")
	s.listening = false

	err := s.conn.Close()
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Broadcast server stopped!")
}

// handlePacket processes the incoming data
func (s *Server) handlePacket(data []byte, clientAddr *net.UDPAddr) {
	body, err := s.codec.DecodeFromBytes(data)
	if err != nil {
		log.Error("Error decoding message framing:", err)
		return
	}

	var msg socket.Message
	err = s.codec.DecodeCBORMessage(body, &msg) // Decode CBOR-encoded body
	if err != nil {
		log.Error("Error decoding CBOR message:", err)
		return
	}

	switch msg.Type {
	case socket.BroadcastMessage:
		log.Infof("Handling message from: %s", clientAddr.String())
		s.handleBroadcast(msg, clientAddr)
	default:
		log.Warn("Unknown message type")
	}
}

func (s *Server) handleBroadcast(msg socket.Message, senderAddr *net.UDPAddr) {
	s.broadcast(msg, func(addr *net.UDPAddr) bool {
		return addr.String() != senderAddr.String()
	})
}

func (s *Server) broadcast(msg socket.Message, eval func(*net.UDPAddr) bool) {
	handlers, found := s.broadcastHandlers[msg.PayloadType]
	if !found {
		return
	}

	s.clientsMux.RLock()
	clientsCopy := make([]net.UDPAddr, len(s.clients))
	for idx, addr := range s.clients {
		clientsCopy[idx] = *addr
	}
	s.clientsMux.RUnlock()

	for _, handler := range handlers {
		response := handler(msg)

		for _, clientAddr := range clientsCopy {
			if !eval(&clientAddr) {
				continue
			}

			if response != nil {
				marshall, err := response.Marshall(s.codec)
				if err != nil {
					return
				}
				_, err = s.conn.WriteToUDP(marshall, &clientAddr)
				if err != nil {
					log.Error("Failed to send message to client:", err)
				}
			}
		}

	}
}
