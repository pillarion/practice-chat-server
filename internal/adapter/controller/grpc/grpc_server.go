package grpc

import (
	"github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
)

type server struct {
	desc.UnimplementedChatV1Server
	chatService chat.Service
}

// NewServer creates a new server.
//
// It takes a service of type chat.Service as a parameter and returns a pointer to server.
func NewServer(cs chat.Service) *server {
	return &server{
		chatService: cs,
	}
}
