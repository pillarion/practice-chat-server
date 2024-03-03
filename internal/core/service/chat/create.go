package chat

import (
	"context"
	"fmt"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	journalModel "github.com/pillarion/practice-chat-server/internal/core/model/journal"
)

func (s *service) CreateChat(ctx context.Context, username []desc.Username) (int64, error) {
	var res int64
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			var errTx error
			res, errTx = s.repoChat.Insert(ctx, username)
			if errTx != nil {
				return errTx
			}

			_, errTx = s.repoJournal.Insert(ctx, &journalModel.Journal{
				Action: "Chat created",
			})
			if errTx != nil {
				return errTx
			}

			return nil
		})
	if err != nil {
		return 0, err
	}
	if res == 0 {
		return 0, fmt.Errorf("failed to create chat")
	}

	return res, nil
}
