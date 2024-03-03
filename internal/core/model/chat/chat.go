package chat

// Username is a model for a username.
type Username string

// Chat is a model for a chat.
type Chat struct {
	ID       int64      `db:"id"`
	Username []Username `db:"chat_name"`
}
