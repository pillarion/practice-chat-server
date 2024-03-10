package message

import (
	"context"

	model "github.com/pillarion/practice-chat-server/internal/core/model/message"
)

// Repo is the interface for the message repository.
//
//go:generate minimock -o mock/ -s "_minimock.go"
type Repo interface {

	// InsertMessage inserts a new message into the database.
	//
	// ctx - the context for the database operation.
	// message - the message object to be inserted.
	// error - any error encountered.
	Insert(ctx context.Context, message *model.Message) error
}
