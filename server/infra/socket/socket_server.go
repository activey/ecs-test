package socket

import (
	"context"
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/codec"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/panjf2000/gnet/v2"
	"sync"
)

type ServerConfig struct {
	multicore bool
	port      int
}

func NewSocketServerConfig() ServerConfig {
	return ServerConfig{
		multicore: true,
		port:      8665,
	}
}

type DisconnectHandlers []func(conn socket.Connection)
type RequestHandlers map[socket.PayloadType]func(writer socket.Connection, request socket.Message) socket.Message
type BroadcastHandlers map[socket.PayloadType][]func(broadcast socket.Message)

type Server struct {
	gnet.BuiltinEventEngine

	requestHandlers    RequestHandlers
	broadcastHandlers  BroadcastHandlers
	disconnectHandlers DisconnectHandlers

	engine    gnet.Engine
	multicore bool
	port      int

	codec *codec.CborCodec
}

func NewServer(config ServerConfig) *Server {
	return &Server{
		multicore:          config.multicore,
		port:               config.port,
		requestHandlers:    make(RequestHandlers),
		broadcastHandlers:  make(BroadcastHandlers),
		disconnectHandlers: make(DisconnectHandlers, 0),

		codec: codec.NewCborCodec(),
	}
}

func (s *Server) OnBoot(engine gnet.Engine) gnet.Action {
	log.Info("Server listening...")
	s.engine = engine
	return gnet.None
}

func (s *Server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Infof("Connected: %s", c.RemoteAddr())
	return
}

func (s *Server) OnClose(c gnet.Conn, err error) gnet.Action {
	log.Infof("Disconnected: %s", c.RemoteAddr())

	cont := c.Context()
	if cont == nil {
		return gnet.None
	}

	sessionId := cont.(session.SessionId)
	if sessionId != "" {
		for _, handler := range s.disconnectHandlers {
			handler(socket.NewDefaultSocketConnection(c, s.codec))
		}
	}

	return gnet.None
}

func (s *Server) OnTraffic(c gnet.Conn) gnet.Action {
	body, err := s.codec.Decode(c)
	if err != nil {
		return 0
	}

	var msg socket.Message
	err = s.codec.DecodeCBORMessage(body, &msg)
	if err != nil {
		log.Error(err)
		return gnet.None
	}

	switch msg.Type {
	case socket.BroadcastMessage:
		//log.Info("Broadcast message received")
		s.handleBroadcast(msg)

	case socket.RequestMessage:
		log.Infof("Request received with ID: %s", msg.RequestID)
		s.handleRequest(c, msg)

	case socket.ResponseMessage:
		log.Info("Response received for RequestID:", msg.RequestID)
		// Handle response for the corresponding request

	default:
		log.Warn("Unknown message type")
	}

	return gnet.None
}

func (s *Server) handleRequest(c gnet.Conn, msg socket.Message) {
	handler, found := s.requestHandlers[msg.PayloadType]
	if found {
		response := handler(socket.NewDefaultSocketConnection(c, s.codec), msg)

		data, err := response.Marshall(s.codec)
		if err != nil {
			log.Error(err)
			return
		}
		_, err = c.Write(data)
		if err != nil {
			log.Error(err)
			return
		}
	} else {
		log.Error("Unsupported PayloadType")
	}
}

func (s *Server) handleBroadcast(msg socket.Message) {
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

func (s *Server) Start() {
	log.Infof("Starting Socket server on port: %d", s.port)

	err := gnet.Run(
		s,
		fmt.Sprintf("tcp://:%d", s.port),
		gnet.WithMulticore(s.multicore),
		gnet.WithReusePort(true),
	)
	if err != nil {
		log.Fatal("An error occurred while starting Socket server", err)
	}
}

func (s *Server) Stop() {
	log.Info("Stopping Socket server...")
	ctx := context.TODO()
	err := s.engine.Stop(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Socket server stopped!")
}

func (s *Server) OnRequest(
	payloadType socket.PayloadType,
	handleFunc func(socket.Connection, socket.Message) socket.Message,
) {
	s.requestHandlers[payloadType] = handleFunc
}

func (s *Server) OnBroadcast(
	payloadType socket.PayloadType,
	handleFunc func(socket.Message),
) {
	s.broadcastHandlers[payloadType] = append(s.broadcastHandlers[payloadType], handleFunc)
}

func (s *Server) OnDisconnect(handler func(conn socket.Connection)) {
	s.disconnectHandlers = append(s.disconnectHandlers, handler)
}
