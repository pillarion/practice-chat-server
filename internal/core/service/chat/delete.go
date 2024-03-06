package chat

import (
	"context"

	journalModel "github.com/pillarion/practice-chat-server/internal/core/model/journal"
)

func (s *service) DeleteChat(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			errTX := s.repoChat.Delete(ctx, id)
			if errTX != nil {
				return errTX
			}

			_, errTX = s.repoJournal.Insert(ctx, &journalModel.Journal{
				Action: "Chat deleted",
			})
			if errTX != nil {
				return errTX
			}

			return nil
		})

	return err
}
