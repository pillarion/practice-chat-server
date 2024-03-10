package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"

	server "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	modelChat "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	chatService "github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	chatServiceMock "github.com/pillarion/practice-chat-server/internal/core/port/service/chat/mock"
)

func TestServer_Create(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) chatService.Service
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()

		usernames = []modelChat.Username{
			modelChat.Username(gofakeit.Username()),
			modelChat.Username(gofakeit.Username()),
			modelChat.Username(gofakeit.Username()),
		}

		req = &desc.CreateRequest{
			Usernames: []string{
				string(usernames[0]),
				string(usernames[1]),
				string(usernames[2]),
			},
		}

		newID = int64(gofakeit.Number(1, 1000))
		res   = &desc.CreateResponse{Id: newID}
	)

	tests := []struct {
		name    string
		service chatServiceMockFunc
		args    args
		want    *desc.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			service: func(mc *minimock.Controller) chatService.Service {
				mock := chatServiceMock.NewServiceMock(mc).
					CreateChatMock.Expect(ctx, usernames).
					Return(newID, nil)

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    res,
			wantErr: false,
		},
		{
			name: "failed case",
			service: func(mc *minimock.Controller) chatService.Service {
				mock := chatServiceMock.NewServiceMock(mc).
					CreateChatMock.Expect(ctx, usernames).
					Return(0, gofakeit.Error())

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := minimock.NewController(t)
			api := server.NewServer(tt.service(mc))

			res, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}

}
