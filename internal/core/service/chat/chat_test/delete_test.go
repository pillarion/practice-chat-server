package chat_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	journalModel "github.com/pillarion/practice-chat-server/internal/core/model/journal"
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
)

func Test_service_DeleteChat(t *testing.T) {
	t.Parallel()
	type chatRepoMockFunc func(mc *minimock.Controller) chatRepo.Repo
	type journalRepoMockFunc func(mc *minimock.Controller) journalRepo.Repo
	type messageRepoMockFunc func(mc *minimock.Controller) messageRepo.Repo
	type txmFunc func(mc *minimock.Controller) txmanager.Transactor
	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()

		txOption  = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		tx        = &pgxMock.Pgtx{}
		ctxWithTx = context.WithValue(ctx, txmanager.TxKey, tx)

		jrnl = &journalModel.Journal{
			Action: "Chat deleted",
		}
		jid = int64(gofakeit.Number(1, 1000))

		ID = int64(gofakeit.Number(1, 1000))
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
				mock := chatRepoMock.NewRepoMock(mc).
					DeleteMock.Expect(ctxWithTx, ID).
					Return(nil)

				return mock
			},
			journalRepoMock: func(mc *minimock.Controller) journalRepo.Repo {
				mock := journalRepoMock.NewRepoMock(mc).
					InsertMock.Expect(ctxWithTx, jrnl).
					Return(jid, nil)

				return mock
			},
			messageRepoMock: func(mc *minimock.Controller) messageRepo.Repo {
				mock := messageRepoMock.NewRepoMock(mc)
				return mock
			},
			transactor: func(mc *minimock.Controller) txmanager.Transactor {
				m := txmanagerMock.NewTransactorMock(mc).
					BeginTxMock.Expect(ctx, txOption).Return(tx, nil)

				return m
			},
			args: args{
				ctx: ctx,
				id:  ID,
			},
			wantErr: false,
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

			err := srvc.DeleteChat(tt.args.ctx, tt.args.id)

			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
