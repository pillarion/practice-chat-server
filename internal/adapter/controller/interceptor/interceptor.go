package interceptor

import (
	accessClient "github.com/pillarion/practice-chat-server/internal/core/tools/access_v1"
)

// ChatServerInterceptor implements grpc.UnaryServerInterceptor
type ChatServerInterceptor struct {
	accessClient accessClient.AccessV1Client
}

// NewChatServerInterceptor creates a new ChatServerInterceptor
func NewChatServerInterceptor(accessClient accessClient.AccessV1Client) *ChatServerInterceptor {
	return &ChatServerInterceptor{
		accessClient: accessClient,
	}
}
