package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	db "github.com/pillarion/practice-platform/pkg/dbclient"
)

// DeleteChat deletes a chat by its ID.
//
// It takes a context.Context and an int64 as parameters and returns an error.
func (p *pg) Delete(ctx context.Context, id int64) error {
	builderDeleteChatUsers := sq.Delete(chatsUsersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatUsersTableChatIDColumn: id})
	query, args, err := builderDeleteChatUsers.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "Chat.Delete",
		QueryRaw: query,
	}
	_, err = p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	builderDeleteChat := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatsTableIDColumn: id})
	query, args, err = builderDeleteChat.ToSql()
	if err != nil {
		return err
	}
	q = db.Query{
		Name:     "Chat.Delete",
		QueryRaw: query,
	}
	_, err = p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
