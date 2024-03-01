package message

import "time"

// Message is a model for a message.
type Message struct {
	From      Username
	Content   string
	Timestamp time.Time
}

// Username is a model for a username.
type Username string

// Chat is a model for a chat.
type Chat struct {
	ID       int64
	Username []Username
}
