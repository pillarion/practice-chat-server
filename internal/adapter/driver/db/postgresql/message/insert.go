package message

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	dto "github.com/pillarion/practice-chat-server/internal/core/dto/postgresql"
	desc "github.com/pillarion/practice-chat-server/internal/core/model/message"
	db "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgclient"
)

// InsertMessage inserts a message into the database.
//
// ctx context.Context, message *desc.Message
// int64, error
func (p *pg) Insert(ctx context.Context, message *desc.Message) error {
	var messageDTO dto.MessageDTO
	if err := messageDTO.FromMessage(message); err != nil {
		return err
	}

	buiderInsert := sq.Insert(messagesTable).
		PlaceholderFormat(sq.Dollar).
		Columns(messagesTableFromColumn, messagesTableContentColumn, messagesTableTimestampColumn).
		Values(messageDTO.From, messageDTO.Content, messageDTO.Timestamp)
	query, args, err := buiderInsert.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "Message.Insert",
		QueryRaw: query,
	}
	_, err = p.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil

}
