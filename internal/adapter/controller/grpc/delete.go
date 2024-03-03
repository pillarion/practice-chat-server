package grpc

import (
	"context"
	"log/slog"

	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete implements desc.ChatV1Server
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.DeleteChat(context.Background(), req.Id)
	if err != nil {
		slog.Error("failed to delete chat", "Error", err)

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
