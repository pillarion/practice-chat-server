package postgresql

import (
	"strconv"
	"time"

	desc "github.com/pillarion/practice-chat-server/internal/core/model/chat"
)

// MessageDTO is a DTO for a desc.Message.
type MessageDTO struct {
	From      int64
	Content   string
	Timestamp time.Time
}

// UsernameDTO is a type alias for a string.
type UsernameDTO string

// FromMessage converts a desc.Message to a MessageDTO.
//
// It takes a pointer to a MessageDTO and a pointer to a desc.Message as parameters.
// It does not return anything.
func (dto *MessageDTO) FromMessage(m *desc.Message) error {
	from, err := strconv.ParseInt(string(m.From), 10, 64)
	if err != nil {
		return err
	}
	dto.From = from
	dto.Content = m.Content
	dto.Timestamp = m.Timestamp

	return nil
}
