package chat

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// SendMessage sends a message using the given context and message.
//
// ctx: context.Context
// message: *desc.Message
// error
func (s *service) SendMessage(ctx context.Context, message *desc.Message) error {

	return s.repoMessage.Insert(ctx, message)
}
