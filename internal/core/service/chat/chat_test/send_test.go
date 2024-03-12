package chat_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	chatModel "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	journalModel "github.com/pillarion/practice-chat-server/internal/core/model/journal"
	messageModel "github.com/pillarion/practice-chat-server/internal/core/model/message"
	chatRepo "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
	chatRepoMock "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat/mock"
	journalRepo "github.com/pillarion/practice-chat-server/internal/core/port/repository/journal"
	journalRepoMock "github.com/pillarion/practice-chat-server/internal/core/port/repository/journal/mock"
	messageRepo "github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
	messageRepoMock "github.com/pillarion/practice-chat-server/internal/core/port/repository/message/mock"
	service "github.com/pillarion/practice-chat-server/internal/core/service/chat"
	pgxMock "github.com/pillarion/practice-platform/pkg/dbclient/mock"
	txmanager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
	txmanagerMock "github.com/pillarion/practice-platform/pkg/pgtxmanager/mock"
	"github.com/stretchr/testify/require"
)

func Test_service_SendMessage(t *testing.T) {
	t.Parallel()
	type chatRepoMockFunc func(mc *minimock.Controller) chatRepo.Repo
	type journalRepoMockFunc func(mc *minimock.Controller) journalRepo.Repo
	type messageRepoMockFunc func(mc *minimock.Controller) messageRepo.Repo
	type txmFunc func(mc *minimock.Controller) txmanager.Transactor

	type args struct {
		ctx     context.Context
		message *messageModel.Message
	}

	var (
		ctx = context.Background()

		txOption  = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		tx        = &pgxMock.Pgtx{}
		ctxWithTx = context.WithValue(ctx, txmanager.TxKey, tx)

		jrnl = &journalModel.Journal{
			Action: "Message sent",
		}
		jid = int64(gofakeit.Number(1, 1000))

		message = &messageModel.Message{
			From:    chatModel.Username(gofakeit.Username()),
			Content: gofakeit.Sentence(10),
		}
	)

	tests := []struct {
		name            string
		chatRepoMock    chatRepoMockFunc
		journalRepoMock journalRepoMockFunc
		messageRepoMock messageRepoMockFunc
		transactor      txmFunc
		args            args
		want            error
		wantErr         bool
	}{
		{
			name: "success case",
			chatRepoMock: func(mc *minimock.Controller) chatRepo.Repo {
				mock := chatRepoMock.NewRepoMock(mc)

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, jrnl).
					Return(jid, nil)

				return mock
			},
			messageRepoMock: func(mc *minimock.Controller) messageRepo.Repo {
				mock := messageRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, message).
					Return(nil)

				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			args: args{
				ctx:     ctx,
				message: message,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "failed case",
			chatRepoMock: func(mc *minimock.Controller) chatRepo.Repo {
				mock := chatRepoMock.NewRepoMock(mc)

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc)

				return mock
			},
			messageRepoMock: func(mc *minimock.Controller) messageRepo.Repo {
				mock := messageRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, message).
					Return(gofakeit.Error())

				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			args: args{
				ctx:     ctx,
				message: message,
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := minimock.NewController(t)
			chatRepo := tt.chatRepoMock(mc)
			journalRepo := tt.journalRepoMock(mc)
			messageRepo := tt.messageRepoMock(mc)
			txManager := txmanager.NewTransactionManager(tt.transactor(mc))

			srvc := service.NewService(chatRepo, messageRepo, journalRepo, txManager)

			err := srvc.SendMessage(tt.args.ctx, tt.args.message)

			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
