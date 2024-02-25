package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// MessageDTO is a DTO for a desc.Message.
type MessageDTO struct {
	From      string
	Content   string
	Timestamp timestamppb.Timestamp
}

// ToMessage converts a MessageDTO to a desc.Message.
func (m *MessageDTO) ToMessage() *desc.Message {
	res := &desc.Message{
		From:      desc.Username(m.From),
		Content:   m.Content,
		Timestamp: m.Timestamp.AsTime(),
	}

	return res
}
