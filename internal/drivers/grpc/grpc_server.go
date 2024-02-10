package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatV1Server
}

// NewServer creates a new server instance.
//
// No parameters. Returns a pointer to a server.
func NewServer() *server {
	return &server{
		UnimplementedChatV1Server: desc.UnimplementedChatV1Server{},
	}
}

// Create implements desc.ChatV1Server
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	slog.Info("GetMessages", "request", req)
	defer ctx.Done()
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// Delete implements desc.ChatV1Server
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("GetMessages", "request", req)
	defer ctx.Done()
	return &emptypb.Empty{}, nil
}

// SendMessage implements desc.ChatV1Server
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	slog.Info("GetMessages", "request", req)
	defer ctx.Done()
	return &emptypb.Empty{}, nil
}
