package grpc

import (
	"context"
	"log/slog"

	dto "github.com/pillarion/practice-chat-server/internal/core/dto/grpc"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage implements desc.ChatV1Server
func (s *Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	messageDTO := dto.MessageDTO{
		From:      req.GetFrom(),
		Content:   req.GetText(),
		Timestamp: *req.GetTimestamp(),
	}

	err := s.chatService.SendMessage(ctx, messageDTO.ToMessage())
	if err != nil {
		slog.Error("failed to send message", "Error", err)

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
