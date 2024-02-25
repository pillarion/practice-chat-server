package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// InsertChat inserts a chat into the database for the given usernames.
//
// ctx context.Context, username *[]desc.Username
// int64, error
func (p *pg) Insert(ctx context.Context, username []desc.Username) (int64, error) {
	userDTO := username
	chatName := "chat:"
	for _, user := range userDTO {
		chatName += fmt.Sprintf(":%s:", user)
	}

	builderInsertChat := sq.Insert(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatsTableNameColumn).
		Values(chatName).
		Suffix("RETURNING " + chatsTableIDColumn)
	query, args, err := builderInsertChat.ToSql()
	if err != nil {

		return 0, err
	}
	var chatID int64
	err = p.pgx.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {

		return 0, err
	}

	builderInsertChatUsers := sq.Insert(chatsUsersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatUsersTableChatIDColumn, chatUsersTableUserIDColumn)
	for _, user := range userDTO {
		builderInsertChatUsers = builderInsertChatUsers.Values(chatID, user)

	}
	query, args, err = builderInsertChatUsers.ToSql()
	if err != nil {

		return 0, err
	}
	_, err = p.pgx.Exec(ctx, query, args...)
	if err != nil {

		return 0, err
	}

	return chatID, nil
}
