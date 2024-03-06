package chat

import (
	"context"

	"github.com/pillarion/practice-chat-server/internal/core/model/chat"
	"github.com/pillarion/practice-chat-server/internal/core/model/message"
)

// Service is the interface for the chat service.
type Service interface {
	CreateChat(ctx context.Context, username []chat.Username) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *message.Message) error
}
