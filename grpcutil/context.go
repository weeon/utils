package grpcutil

import (
	"context"
	"github.com/weeon/contract"
	"github.com/weeon/utils/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)


func RequestIDFromIncomingContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if v, ok := md[contract.RequestID]; ok {
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
		contract.RequestID: reqID,
	})
	ctx = metadata.NewIncomingContext(ctx, md)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

// Wrap Request ID from incoming context
func UnaryServerWrapRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := ctxutil.AddRequestID(ctx,RequestIDFromIncomingContext(ctx))
		resp, err := handler(newCtx, req)
		return resp, err
	}
}

