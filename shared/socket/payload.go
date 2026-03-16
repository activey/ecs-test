package socket

type Payload interface {
	Type() PayloadType
	DecodeFrom(data []byte) error
	Encode() ([]byte, error)
}
