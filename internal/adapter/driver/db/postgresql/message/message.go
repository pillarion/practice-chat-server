package message

import (
	db "github.com/pillarion/practice-platform/pkg/dbclient"

	"github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
)

const (
	messagesTable                = "messages"
	messagesTableIDColumn        = "id"
	messagesTableFromColumn      = "from_id"
	messagesTableContentColumn   = "content"
	messagesTableTimestampColumn = "timestamp"
)

type pg struct {
	db db.Client
}

// New initializes a new user repository using the provided database configuration.
//
// ctx context.Context, cfg *config.Database
// repo.ChatRepo, error
func New(db db.Client) (message.Repo, error) {

	return &pg{
		db: db,
	}, nil
}
