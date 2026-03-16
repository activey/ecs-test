package client

import (
	"ecs-test/client/config"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/codec"
	"ecs-test/shared/socket/payload"
	"fmt"
	"github.com/charmbracelet/log"
	"sync"
	"time"

	"net"
)

type Handlers map[socket.PayloadType][]func(broadcast socket.Message)

type ServerBroadcastClient struct {
	conn              *net.UDPConn
	serverAddr        *net.UDPAddr
	codec             *codec.CborCodec
	broadcastHandlers Handlers

	connected bool
}

func NewServerBroadcastClient(cfg config.GameClientConfig) *ServerBroadcastClient {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", cfg.ServerAddress, 8666))
	if err != nil {
		return nil
	}

	return &ServerBroadcastClient{
		serverAddr:        addr,
		codec:             codec.NewCborCodec(),
		broadcastHandlers: make(Handlers),
	}
}

func (s *ServerBroadcastClient) Connect() error {
	log.Infof("Connecting to Broadcast server at: %s", s.serverAddr)
	conn, err := net.DialUDP("udp", nil, s.serverAddr)
	if err != nil {
		return err
	}
	s.conn = conn
	log.Info("Connected to server")

	s.connected = true
	return nil
}

func (s *ServerBroadcastClient) Start() {
	go s.listen()
	go s.broadcastPing()
}

func (s *ServerBroadcastClient) listen() {
	buffer := make([]byte, 1024)
	for s.connected {
		n, clientAddr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Error("Failed to read from UDP:", err)
			continue
		}
		go s.handlePacket(buffer[:n], clientAddr)
	}
}

func (s *ServerBroadcastClient) broadcastPing() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C: // Wait for the ticker to send a ping every second
			if !s.connected {
				return // Exit the loop if not connected
			}
			err := s.SendBroadcast(payload.NewPingBroadcast())
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (s *ServerBroadcastClient) handlePacket(data []byte, addr *net.UDPAddr) {
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
		s.handleBroadcast(msg)
	default:
		log.Warn("Unknown message type")
	}
}

func (s *ServerBroadcastClient) handleBroadcast(msg socket.Message) {
	handlers, found := s.broadcastHandlers[msg.PayloadType]
	if found {
		var wg sync.WaitGroup
		for _, handler := range handlers {
			wg.Add(1)
			go func(handleFunc func(msg socket.Message)) {
				defer wg.Done()
				handleFunc(msg)
			}(handler)
		}
		wg.Wait()
	} else {
		log.Error("Unsupported PayloadType")
	}
}

func (s *ServerBroadcastClient) OnBroadcast(
	payloadType socket.PayloadType,
	handleFunc func(socket.Message),
) {
	s.broadcastHandlers[payloadType] = append(s.broadcastHandlers[payloadType], handleFunc)
}

func (s *ServerBroadcastClient) SendBroadcast(payload socket.Payload) error {
	if !s.connected {
		return nil
	}
	broadcastMessage, err := socket.NewBroadcast(payload)
	if err != nil {
		return err
	}

	data, err := broadcastMessage.Marshall(s.codec)
	if err != nil {
		return err
	}

	_, err = s.conn.Write(data)
	return err
}

func (s *ServerBroadcastClient) Disconnect() {
	s.connected = false

	err := s.conn.Close()
	if err != nil {
		log.Error(err)
	}
}
