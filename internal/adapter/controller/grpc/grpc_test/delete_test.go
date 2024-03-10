package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	server "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	chatService "github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	chatServiceMock "github.com/pillarion/practice-chat-server/internal/core/port/service/chat/mock"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
)

func TestServer_Delete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) chatService.Service

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()

		ID = int64(gofakeit.Number(1, 1000))

		req = &desc.DeleteRequest{
			Id: ID,
		}
	)

	tests := []struct {
		name    string
		service chatServiceMockFunc
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			service: func(mc *minimock.Controller) chatService.Service {
				mock := chatServiceMock.NewServiceMock(mc).
					DeleteChatMock.Expect(ctx, ID).
					Return(nil)

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
		{
			name: "failed case",
			service: func(mc *minimock.Controller) chatService.Service {
				mock := chatServiceMock.NewServiceMock(mc).
					DeleteChatMock.Expect(ctx, ID).
					Return(gofakeit.Error())

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

			res, err := api.Delete(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
