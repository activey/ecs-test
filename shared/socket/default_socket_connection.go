package socket

import (
	"ecs-test/shared/session"
	"ecs-test/shared/socket/codec"
	"github.com/panjf2000/gnet/v2"
)

type DefaultSocketConnection struct {
	conn  gnet.Conn
	codec *codec.CborCodec
}

func NewDefaultSocketConnection(conn gnet.Conn, codec *codec.CborCodec) *DefaultSocketConnection {
	return &DefaultSocketConnection{
		conn:  conn,
		codec: codec,
	}
}

func (a *DefaultSocketConnection) SetSessionId(id session.SessionId) {
	// marking connection as related with given session
	a.conn.SetContext(id)
}

func (a *DefaultSocketConnection) SessionId() session.SessionId {
	return a.conn.Context().(session.SessionId)
}

func (a *DefaultSocketConnection) Write(message Message) error {
	encoded, err := message.Marshall(a.codec)
	if err != nil {
		return err
	}
	return a.conn.AsyncWrite(encoded, nil)
}
