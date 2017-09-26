package msg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

const hdrLength = 12

// MSG defines the message protocol data.
type MSG struct {
	Name string
	Data string
}

// String implements the fmt Stringer interface.
func (m MSG) String() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("\nName: %s\n", m.Name))
	b.WriteString(fmt.Sprintf("Data: %s\n", m.Data))

	return b.String()
}

// Read waits on the network to receive a chat message.
func Read(r io.Reader) ([]byte, int, error) {

	// Read the first header length of bytes.
	buf := make([]byte, hdrLength)
	if _, err := io.ReadFull(r, buf); err != nil {
		errors.Wrap(err, "ReadFull header")
		return nil, 0, err
	}

	// Get the length for the remaining bytes.
	length := int(binary.BigEndian.Uint16(buf[10:12])) + hdrLength

	// Copy the header bytes into the final slice.
	data := make([]byte, length)
	copy(data, buf)

	// Read the remaining bytes.
	if _, err := io.ReadFull(r, data[hdrLength:]); err != nil {
		errors.Wrap(err, "ReadFull data")
		return nil, 0, err
	}

	return data, length, nil
}

// Decode will take the bytes and create a MSG value.
func Decode(data []byte) MSG {
	var name string
	if n := bytes.IndexByte(data[:10], 0); n != -1 {
		name = string(data[:n])
	} else {
		name = string(data[:10])
	}
	return MSG{
		Name: name,
		Data: string(data[12:]),
	}
}

// Encode will take a message and produce byte slice.
func Encode(msg MSG) []byte {
	// we can't have more than the first 10 bytes.
	n := len(msg.Name)
	if n > 10 {
		n = 10
	}
	data := make([]byte, hdrLength+len(msg.Data))

	copy(data, msg.Name[:n])
	binary.BigEndian.PutUint16(data[10:12], uint16(len(msg.Data)))
	copy(data[12:], msg.Data)

	return data
}
