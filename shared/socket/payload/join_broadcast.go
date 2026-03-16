package payload

import (
	"ecs-test/shared/player"
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"github.com/fxamacker/cbor/v2"
)

type JoinBroadcast struct {
	SessionId session.SessionId `cbor:"sid"`
	UserName  string            `cbor:"u"`
	Position  player.Position   `cbor:"p"`
}

func NewJoinBroadcast() *JoinBroadcast {
	return &JoinBroadcast{}
}

func (b *JoinBroadcast) WithSessionId(id session.SessionId) *JoinBroadcast {
	b.SessionId = id
	return b
}

func (b *JoinBroadcast) WithWithUserName(userName string) *JoinBroadcast {
	b.UserName = userName
	return b
}

func (b *JoinBroadcast) WithPosition(position player.Position) *JoinBroadcast {
	b.Position = position
	return b
}

func (b *JoinBroadcast) Type() socket.PayloadType {
	return socket.PlayerJoinBroadcast
}

func (b *JoinBroadcast) DecodeFrom(data []byte) error {
	return cbor.Unmarshal(data, b)
}

func (b *JoinBroadcast) Encode() ([]byte, error) {
	return cbor.Marshal(b)
}
