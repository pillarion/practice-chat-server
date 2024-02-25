package chat

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// Repo is the interface for the chat repository.
type Repo interface {
	// InsertUser inserts a new user into the database.
	//
	// ctx - the context for the database operation.
	// user - the user object to be inserted.
	// (int64, error) - returns the user_id of the inserted user and any error encountered.
	Insert(ctx context.Context, username []desc.Username) (int64, error)

	// DeleteChat deletes a chat from the database.
	//
	// ctx - the context for the database operation.
	// id - the id of the chat to be deleted.
	// error - any error encountered.
	Delete(ctx context.Context, id int64) error
}
