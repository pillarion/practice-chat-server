package grpc

import (
	"context"

	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete implements desc.ChatV1Server
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.DeleteChat(context.Background(), req.Id)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
