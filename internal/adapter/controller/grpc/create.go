package grpc

import (
	"context"
	"log/slog"

	model "github.com/pillarion/practice-chat-server/internal/core/model/chat"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
)

// Create implements desc.ChatV1Server
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	usernames := make([]model.Username, len(req.Usernames))
	for i, username := range req.Usernames {
		usernames[i] = model.Username(username)
	}

	id, err := s.chatService.CreateChat(ctx, usernames)
	if err != nil {
		slog.Error("failed to create chat", "Error", err)

		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
