package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// DeleteChat deletes a chat by its ID.
//
// It takes a context.Context and an int64 as parameters and returns an error.
func (p *pg) Delete(ctx context.Context, id int64) error {

	builderDeleteChatUsers := sq.Delete("chats_users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": id})

	query, args, err := builderDeleteChatUsers.ToSql()
	if err != nil {
		return err
	}

	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	builderDeleteChat := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err = builderDeleteChat.ToSql()
	if err != nil {
		return err
	}

	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
