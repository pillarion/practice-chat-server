package chat

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
)

const (
	chatNamePrefix = "chat:"

	chatsTable           = "chats"
	chatsTableIDColumn   = "id"
	chatsTableNameColumn = "name"

	chatsUsersTable            = "chats_users"
	chatUsersTableChatIDColumn = "chat_id"
	chatUsersTableUserIDColumn = "user_id"
)

type pg struct {
	pgx *pgxpool.Pool
}

// New initializes a new user repository using the provided database configuration.
//
// ctx context.Context, cfg *config.Database
// repo.ChatRepo, error
func New(pgx *pgxpool.Pool) (chat.Repo, error) {

	return &pg{
		pgx: pgx,
	}, nil
}
