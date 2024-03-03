package journal

import (
	"github.com/pillarion/practice-chat-server/internal/core/port/repository/journal"
	db "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgclient"
)

const (
	journalTable                = "journal"
	journalTableIDColumn        = "id"
	journalTableActionColumn    = "action"
	journalTableCreatedAtColumn = "created_at"
)

type pg struct {
	db db.Client
}

// New initializes a new user repository using the provided database configuration.
//
// db: the database client.
// repo.UserRepo, error
func New(db db.Client) (journal.Repo, error) {
	return &pg{
		db: db,
	}, nil
}
