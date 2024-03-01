package chat

import (
	"context"
)

func (s *service) DeleteChat(ctx context.Context, id int64) error {
	return s.repoChat.Delete(ctx, id)
}
