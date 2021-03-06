package reader

import (
	"time"

	"github.com/elastic/beats/libbeat/common"
)

// Message represents a reader event with timestamp, content and actual number
// of bytes read from input before decoding.
type Message struct {
	Ts      time.Time     // timestamp the content was read
	Content []byte        // actual content read
	Bytes   int           // total number of bytes read to generate the message
	Fields  common.MapStr // optional fields that can be added by reader
}

// Reader is the interface that wraps the basic Next method for
// getting a new message.
// Next returns the message being read or and error. EOF is returned
// if reader will not return any new message on subsequent calls.
type Reader interface {
	Next() (Message, error)
}

// IsEmpty returns true in case the message is empty
// A message with only newline character is counted as an empty message
func (m *Message) IsEmpty() bool {
	// If no Bytes were read, event is empty
	// For empty line Bytes is at least 1 because of the newline char
	if m.Bytes == 0 {
		return true
	}

	// Content length can be 0 because of JSON events. Content and Fields must be empty.
	if len(m.Content) == 0 && len(m.Fields) == 0 {
		return true
	}

	return false
}
