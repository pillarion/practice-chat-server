package chat

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// Service is the interface for the chat service.
type Service interface {
	CreateChat(ctx context.Context, username []desc.Username) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *desc.Message) error
}
