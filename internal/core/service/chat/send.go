package chat

import (
	"context"

	journalModel "github.com/pillarion/practice-chat-server/internal/core/model/journal"
	desc "github.com/pillarion/practice-chat-server/internal/core/model/message"
)

// SendMessage sends a message using the given context and message.
//
// ctx: context.Context
// message: *desc.Message
// error
func (s *service) SendMessage(ctx context.Context, message *desc.Message) error {
	err := s.txManager.ReadCommitted(
		ctx,
		func(ctx context.Context) error {
			errTX := s.repoMessage.Insert(ctx, message)
			if errTX != nil {
				return errTX
			}

			_, errTX = s.repoJournal.Insert(ctx, &journalModel.Journal{
				Action: "Message sent",
			})
			if errTX != nil {
				return errTX
			}

			return nil
		})

	return err
}
