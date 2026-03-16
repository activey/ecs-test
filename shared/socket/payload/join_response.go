package payload

import (
	"ecs-test/shared/player"
	"ecs-test/shared/socket"
	"github.com/fxamacker/cbor/v2"
	"time"
)

type JoinResponseStatus int

const (
	JoinSuccessful JoinResponseStatus = iota
	JoinFailedAlreadyThere
	JoinFailedUnknownSession
	JoinFailedSystemError
)

type JoinResponse struct {
	Status       JoinResponseStatus `cbor:"s"`
	Time         time.Time          `cbor:"t"`
	Position     player.Position    `cbor:"p"`
	OtherPlayers []player.Player    `cbor:"op"`
}

func NewPlayerJoinResponsePayload() *JoinResponse {
	return &JoinResponse{
		Time:         time.Now(),
		Status:       JoinSuccessful,
		OtherPlayers: make([]player.Player, 0),
	}
}

func (r *JoinResponse) WithStatus(s JoinResponseStatus) *JoinResponse {
	r.Status = s
	return r
}

func (r *JoinResponse) WithOtherPlayers(others ...player.Player) *JoinResponse {
	for _, other := range others {
		r.OtherPlayers = append(r.OtherPlayers, other)
	}
	return r
}

func (r *JoinResponse) WithPosition(position player.Position) *JoinResponse {
	r.Position = position
	return r
}

func (r *JoinResponse) Type() socket.PayloadType {
	return socket.PlayerJoinResponse
}

func (r *JoinResponse) DecodeFrom(data []byte) error {
	return cbor.Unmarshal(data, r)
}

func (r *JoinResponse) Encode() ([]byte, error) {
	return cbor.Marshal(r)
}
