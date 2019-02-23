package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	RequestID = "request-id"
)

func RequestIDFromIncomingContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if v, ok := md[RequestID]; ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

func ContextAddRequestID(ctx context.Context, reqID string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		md[reqID] = []string{reqID}
		ctx = metadata.NewIncomingContext(ctx, md)
		ctx = metadata.NewOutgoingContext(ctx, md)
		return ctx
	}
	md = metadata.New(map[string]string{
		RequestID: reqID,
	})
	ctx = metadata.NewIncomingContext(ctx, md)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}
