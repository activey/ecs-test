package payload

import (
	"ecs-test/shared/session"
	"ecs-test/shared/socket"
	"github.com/fxamacker/cbor/v2"
)

type JoinRequest struct {
	SessionId session.SessionId `cbor:"sid"`
}

func NewPlayerJoinRequest(sessionId session.SessionId) *JoinRequest {
	return &JoinRequest{
		SessionId: sessionId,
	}
}

func (r *JoinRequest) Type() socket.PayloadType {
	return socket.PlayerJoinRequest
}

func (r *JoinRequest) DecodeFrom(data []byte) error {
	return cbor.Unmarshal(data, r)
}

func (r *JoinRequest) Encode() ([]byte, error) {
	return cbor.Marshal(r)
}
