package journal

import (
	"context"

	"github.com/pillarion/practice-chat-server/internal/core/model/journal"
)

// Repo defines the journal repository
type Repo interface {
	Insert(ctx context.Context, j *journal.Journal) (int64, error)
}
