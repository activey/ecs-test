package payload

import (
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"github.com/fxamacker/cbor/v2"
	"time"
)

type PingBroadcast struct {
	Time      time.Time         `cbor:"t"`
	SessionId session.SessionId `cbor:"sid"`
}

func (p *PingBroadcast) Type() socket.PayloadType {
	return socket.PlayerPing
}

func (p *PingBroadcast) DecodeFrom(data []byte) error {
	return cbor.Unmarshal(data, p)
}

func (p *PingBroadcast) Encode() ([]byte, error) {
	return cbor.Marshal(p)
}

func NewPingBroadcast() *PingBroadcast {
	return &PingBroadcast{
		Time: time.Now(),
	}
}

func (p *PingBroadcast) WithSessionId(id session.SessionId) *PingBroadcast {
	p.SessionId = id
	return p
}
