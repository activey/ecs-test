package payload

import (
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"github.com/fxamacker/cbor/v2"
	"time"
)

type PlayerDirection int

type PlayerMovementBroadcast struct {
	Direction int               `cbor:"d"`
	X         float64           `cbor:"x"`
	Y         float64           `cbor:"y"`
	Time      time.Time         `cbor:"t"`
	SessionId session.SessionId `cbor:"s"`
}

func NewPlayerMovementBroadcast() *PlayerMovementBroadcast {
	return &PlayerMovementBroadcast{
		Time: time.Now(),
	}
}

func (p *PlayerMovementBroadcast) WithSessionId(id session.SessionId) *PlayerMovementBroadcast {
	p.SessionId = id
	return p
}

func (p *PlayerMovementBroadcast) Type() socket.PayloadType {
	return socket.PlayerMovementBroadcast
}

func (p *PlayerMovementBroadcast) DecodeFrom(data []byte) error {
	return cbor.Unmarshal(data, p)
}

func (p *PlayerMovementBroadcast) Encode() ([]byte, error) {
	return cbor.Marshal(p)
}

func (p *PlayerMovementBroadcast) WithMovementData(x float64, y float64, direction int) *PlayerMovementBroadcast {
	p.X = x
	p.Y = y
	p.Direction = direction
	return p
}
