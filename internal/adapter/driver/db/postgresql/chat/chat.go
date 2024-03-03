package chat

import (
	db "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgclient"

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
	db db.Client
}

// New initializes a new user repository using the provided database configuration.
//
// db: the database client.
// repo.ChatRepo, error
func New(db db.Client) (chat.Repo, error) {

	return &pg{
		db: db,
	}, nil
}
