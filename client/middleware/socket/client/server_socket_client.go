package client

import (
	"ecs-test/client/config"
	"ecs-test/shared/socket"
	"ecs-test/shared/socket/codec"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/panjf2000/gnet/v2"
)

type RequestPair struct {
	RequestID  socket.RequestId
	ResponseCh chan []byte
}

type RequestResponse map[socket.RequestId]chan []byte

type DisconnectHandlers []func(err error)
type BroadcastHandlers map[socket.PayloadType]func(broadcast socket.Message)

type ServerSocketClient struct {
	gnet.BuiltinEventEngine

	serverAddress      string
	conn               gnet.Conn
	broadcastHandlers  BroadcastHandlers
	disconnectHandlers DisconnectHandlers

	requestResponse RequestResponse
	requestQueue    chan RequestPair
	codec           *codec.CborCodec
}

func NewServerSocketClient(config config.GameClientConfig) *ServerSocketClient {
	return &ServerSocketClient{
		serverAddress:      config.ServerAddress,
		requestResponse:    make(RequestResponse),
		broadcastHandlers:  make(BroadcastHandlers),
		disconnectHandlers: make(DisconnectHandlers, 0),
		requestQueue:       make(chan RequestPair),
		codec:              codec.NewCborCodec(),
	}
}

func (s *ServerSocketClient) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	s.conn = c
	return
}

func (s *ServerSocketClient) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	for _, handler := range s.disconnectHandlers {
		handler(err)
	}
	return
}

func (s *ServerSocketClient) OnTraffic(c gnet.Conn) (action gnet.Action) {
	body, err := s.codec.Decode(c)
	if err != nil {
		return
	}

	var msg socket.Message
	err = s.codec.DecodeCBORMessage(body, &msg)
	if err != nil {
		log.Error(err)
		return
	}

	switch msg.Type {
	case socket.BroadcastMessage:
		s.handlerBroadcast(msg)
	case socket.ResponseMessage:
		if responseChan, ok := s.requestResponse[msg.RequestID]; ok {
			delete(s.requestResponse, msg.RequestID)
			responseChan <- msg.Payload
		}
	default:
		fmt.Println("Unsupported message type, for now ;P")
		// do nothing for now
	}

	return
}

func (s *ServerSocketClient) handlerBroadcast(msg socket.Message) {
	handler, found := s.broadcastHandlers[msg.PayloadType]
	if found {
		handler(msg)
	} else {
		log.Error("Unsupported PayloadType")
	}
}

func (s *ServerSocketClient) SendRequest(payload socket.Payload) (chan []byte, error) {
	responseCh := make(chan []byte)
	request, err := socket.NewRequest(payload)
	if err != nil {
		return nil, err
	}

	data, err := request.Marshall(s.codec)
	if err != nil {
		return nil, err
	}

	err = s.conn.AsyncWrite(data, nil)
	if err != nil {
		return nil, err
	}

	s.requestQueue <- RequestPair{
		RequestID:  request.RequestID,
		ResponseCh: responseCh,
	}
	return responseCh, nil
}

func (s *ServerSocketClient) SendBroadcast(payload socket.Payload) error {
	broadcastMessage, err := socket.NewBroadcast(payload)
	if err != nil {
		return err
	}

	data, err := broadcastMessage.Marshall(s.codec)
	if err != nil {
		return err
	}

	err = s.conn.AsyncWrite(data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerSocketClient) Connect() error {
	log.Info("Starting Socket client ...")

	go s.processRequestResponse()

	cli, err := gnet.NewClient(
		s,
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithLockOSThread(true),
	)

	if err != nil {
		return err
	}

	err = cli.Start()
	if err != nil {
		return err
	}

	c, err := cli.DialContext("tcp", fmt.Sprintf("%s:8665", s.serverAddress), nil)
	if err != nil {
		return err
	}

	err = c.Wake(nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerSocketClient) processRequestResponse() {
	for {
		select {
		case req := <-s.requestQueue:
			s.requestResponse[req.RequestID] = req.ResponseCh
		}
	}
}

func (s *ServerSocketClient) OnBroadcast(
	payloadType socket.PayloadType,
	handleFunc func(broadcast socket.Message),
) {
	s.broadcastHandlers[payloadType] = handleFunc
}

func (s *ServerSocketClient) OnConnectionClosed(handleFunc func(error)) {
	s.disconnectHandlers = append(s.disconnectHandlers, handleFunc)
}

func (s *ServerSocketClient) Disconnect() error {
	return s.conn.Close()
}
