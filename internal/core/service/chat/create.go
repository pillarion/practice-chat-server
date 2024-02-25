package chat

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

func (s *service) CreateChat(ctx context.Context, username []desc.Username) (int64, error) {
	return s.repoChat.Insert(ctx, username)
}
