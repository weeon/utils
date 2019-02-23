package grpc

import (
	"context"
	"testing"
)

func TestGetRequestID(t *testing.T){
	reqID := "abc"
	ctx := context.Background()
	ctx = ContextAddRequestID(ctx,reqID)

	t.Log(RequestIDFromMetadataContext(ctx))
}
