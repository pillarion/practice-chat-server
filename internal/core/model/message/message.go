package message

import (
	"time"

	"github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// Message is a model for a message.
type Message struct {
	From      chat.Username `db:"from_id"`
	Content   string        `db:"content"`
	Timestamp time.Time     `db:"timestamp"`
}
