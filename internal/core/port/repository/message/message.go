package message

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// Repo is the interface for the message repository.
type Repo interface {

	// InsertMessage inserts a new message into the database.
	//
	// ctx - the context for the database operation.
	// message - the message object to be inserted.
	// error - any error encountered.
	Insert(ctx context.Context, message *desc.Message) error
}
