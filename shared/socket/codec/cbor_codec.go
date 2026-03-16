package codec

import (
	"encoding/binary"
	"errors"
	"github.com/fxamacker/cbor/v2"
	"github.com/panjf2000/gnet/v2"
)

var ErrIncompletePacket = errors.New("incomplete packet")

const (
	bodySize = 4 // Size of the length prefix (4 bytes for a 32-bit unsigned int)
)

// CborCodec with length prefix
type CborCodec struct{}

// NewCborCodec returns a new instance of the codec
func NewCborCodec() *CborCodec {
	return &CborCodec{}
}

// Encode encodes the given payload using CBOR and adds a length prefix.
func (codec *CborCodec) Encode(msg interface{}) ([]byte, error) {
	// Encode the message with CBOR
	cborData, err := cbor.Marshal(msg)
	if err != nil {
		return nil, err
	}

	// Create the final packet with a length prefix
	msgLen := bodySize + len(cborData) // total message length: bodySize + CBOR-encoded message
	data := make([]byte, msgLen)

	// Write the length of the CBOR-encoded data (length prefix)
	binary.BigEndian.PutUint32(data[:bodySize], uint32(len(cborData)))

	// Copy the CBOR-encoded data into the buffer after the length prefix
	copy(data[bodySize:], cborData)

	return data, nil
}

// Decode reads and decodes CBOR-encoded data from the connection.
func (codec *CborCodec) Decode(c gnet.Conn) ([]byte, error) {
	// Read the length prefix (4 bytes)
	buf, _ := c.Peek(bodySize)
	if len(buf) < bodySize {
		return nil, ErrIncompletePacket // Not enough data for the length prefix
	}

	// Read the body length (total length of the CBOR-encoded message)
	bodyLen := binary.BigEndian.Uint32(buf[:bodySize])

	// Check if we have enough data buffered for the full message
	if c.InboundBuffered() < int(bodySize+bodyLen) {
		return nil, ErrIncompletePacket // Not enough data for the full message
	}

	// Read the full message (length prefix + CBOR body)
	buf, _ = c.Peek(int(bodySize + bodyLen))
	_, _ = c.Discard(int(bodySize + bodyLen)) // Discard the data we just read

	// Return the body (CBOR-encoded data)
	return buf[bodySize:], nil
}

// DecodeFromBytes decodes a CBOR message from a byte slice (raw input)
func (codec *CborCodec) DecodeFromBytes(data []byte) ([]byte, error) {
	if len(data) < bodySize {
		return nil, ErrIncompletePacket
	}

	bodyLen := binary.BigEndian.Uint32(data[:bodySize])
	if len(data) < int(bodySize+bodyLen) {
		return nil, ErrIncompletePacket
	}

	return data[bodySize : bodySize+bodyLen], nil
}

// DecodeCBORMessage decodes the CBOR message from the raw body bytes.
func (codec *CborCodec) DecodeCBORMessage(body []byte, out interface{}) error {
	return cbor.Unmarshal(body, out)
}
