package chat

import (
	"context"

	modelChat "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	modelMessage "github.com/pillarion/practice-chat-server/internal/core/model/message"
)

// Service is the interface for the chat service.
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Service interface {
	CreateChat(ctx context.Context, username []modelChat.Username) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *modelMessage.Message) error
}
