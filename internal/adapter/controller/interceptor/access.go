package interceptor

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

// VerifyAccessInterceptor is a ChatServerInterceptor implements grpc.UnaryServerInterceptor
func (i *ChatServerInterceptor) VerifyAccessInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, fmt.Errorf("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	md = metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := i.accessClient.Check(
		ctx,
		info.FullMethod,
	)
	if err != nil {
		return nil, err
	}

	// Call the next handler
	return handler(ctx, req)
}
