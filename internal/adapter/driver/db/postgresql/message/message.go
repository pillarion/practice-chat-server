package message

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
)

type pg struct {
	pgx *pgxpool.Pool
}

// New initializes a new user repository using the provided database configuration.
//
// ctx context.Context, cfg *config.Database
// repo.ChatRepo, error
func New(pgx *pgxpool.Pool) (message.Repo, error) {

	return &pg{
		pgx: pgx,
	}, nil
}
