package grpcutil

import (
	"context"
	"testing"
)

func TestGetRequestID(t *testing.T) {
	reqID := "abc"
	ctx := context.Background()
	ctx = ContextAddRequestID(ctx, reqID)
	if RequestIDFromIncomingContext(ctx) != reqID {
		t.Error("cant get request id from conetxt")
	}
}
