package grpc_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	server "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	"github.com/pillarion/practice-chat-server/internal/core/model/chat"
	"github.com/pillarion/practice-chat-server/internal/core/model/message"
	chatService "github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	chatServiceMock "github.com/pillarion/practice-chat-server/internal/core/port/service/chat/mock"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestServer_SendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) chatService.Service

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()

		req = &desc.SendMessageRequest{
			From:      gofakeit.Username(),
			Text:      gofakeit.Sentence(10),
			Timestamp: timestamppb.New(gofakeit.Date()),
		}

		message = &message.Message{
			From:      chat.Username(req.From),
			Content:   req.Text,
			Timestamp: req.Timestamp.AsTime(),
		}

		res = &emptypb.Empty{}
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
					SendMessageMock.Expect(ctx, message).
					Return(nil)

				return mock
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want:    res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := minimock.NewController(t)
			api := server.NewServer(tt.service(mc))

			res, err := api.SendMessage(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
