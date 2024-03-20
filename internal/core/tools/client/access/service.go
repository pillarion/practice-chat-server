package access

import (
	"context"

	accessClient "github.com/pillarion/practice-chat-server/internal/core/tools/access_v1"
)

// V1Service implements access_v1.AccessV1Server
type v1Client struct {
	accessClient accessClient.AccessV1Client
}

// NewV1Client creates a new V1Service
func NewV1Client(accessClient accessClient.AccessV1Client) *v1Client {
	return &v1Client{
		accessClient: accessClient,
	}
}

// Check implements access_v1.AccessV1Server
func (s *v1Client) Check(ctx context.Context, endpoint string) error {
	req := &accessClient.CheckRequest{
		EndpointAddress: endpoint,
	}
	_, err := s.accessClient.Check(ctx, req)

	return err
}
