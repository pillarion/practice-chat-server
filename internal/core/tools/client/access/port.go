package access

import "context"

// V1Client is the interface for the access client
type V1Client interface {
	Check(ctx context.Context, endpoint string) error
}
