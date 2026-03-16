package socket

import (
	"ecs-test/shared/socket/codec"
	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"
)

type MessageType int

const (
	BroadcastMessage MessageType = iota
	RequestMessage
	ResponseMessage
)

type PayloadType int

const (
	PlayerJoinRequest PayloadType = iota
	PlayerJoinResponse

	PlayerJoinBroadcast
	PlayerMovementBroadcast
	PlayerPing
)

type RequestId string

func NewRequestId() RequestId {
	return RequestId(uuid.New().String())
}

type Message struct {
	Type        MessageType `cbor:"t"`  // Message type (broadcast, request, response)
	RequestID   RequestId   `cbor:"id"` // Unique request ID (for request-reply pattern)
	Payload     []byte      `cbor:"p"`  // Encoded payload data (raw bytes)
	PayloadType PayloadType `cbor:"pt"` // The type of payload (PlayerJoinRequest, PlayerMoved, etc.)
}

func (m Message) Marshall(codec *codec.CborCodec) ([]byte, error) {
	return codec.Encode(m)
}

func (m Message) DecodePayload(out interface{}) error {
	err := cbor.Unmarshal(m.Payload, out)
	if err != nil {
		return err
	}
	return nil
}

func NewBroadcast(payload Payload) (Message, error) {
	encodedPayload, err := payload.Encode()
	if err != nil {
		return Message{}, err
	}
	return Message{
		Type:        BroadcastMessage,
		Payload:     encodedPayload,
		PayloadType: payload.Type(),
	}, nil
}

func NewRequest(payload Payload) (Message, error) {
	encodedPayload, err := payload.Encode()
	if err != nil {
		return Message{}, err
	}
	return Message{
		Type:        RequestMessage,
		RequestID:   NewRequestId(),
		Payload:     encodedPayload,
		PayloadType: payload.Type(),
	}, nil
}

func NewResponse(requestID RequestId, payloadType PayloadType, payload Payload) (Message, error) {
	encodedPayload, err := payload.Encode()
	if err != nil {
		return Message{}, err
	}
	return Message{
		Type:        ResponseMessage,
		RequestID:   requestID,
		Payload:     encodedPayload,
		PayloadType: payloadType,
	}, nil
}
