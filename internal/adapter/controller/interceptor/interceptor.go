package interceptor

import (
	accessClient "github.com/pillarion/practice-chat-server/internal/core/tools/client/access"
)

// ChatServerInterceptor implements grpc.UnaryServerInterceptor
type ChatServerInterceptor struct {
	accessClient accessClient.V1Client
}

// NewChatServerInterceptor creates a new ChatServerInterceptor
func NewChatServerInterceptor(accessClient accessClient.V1Client) *ChatServerInterceptor {
	return &ChatServerInterceptor{
		accessClient: accessClient,
	}
}
