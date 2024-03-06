package grpc

import (
	"github.com/pillarion/practice-chat-server/internal/core/model/chat"
	"github.com/pillarion/practice-chat-server/internal/core/model/message"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MessageDTO is a DTO for a desc.Message.
type MessageDTO struct {
	From      string
	Content   string
	Timestamp timestamppb.Timestamp
}

// ToMessage converts a MessageDTO to a desc.Message.
func (m *MessageDTO) ToMessage() *message.Message {
	res := &message.Message{
		From:      chat.Username(m.From),
		Content:   m.Content,
		Timestamp: m.Timestamp.AsTime(),
	}

	return res
}
