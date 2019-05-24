package grpcutil

import (
	"context"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/weeon/contract"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ListenerFromEnv() (net.Listener, error) {
	addr := os.Getenv("SRV_ADDR")
	if len(addr) == 0 {
		addr = ":8282"
	}
	return net.Listen("tcp", addr)
}

func NewServer(logger *zap.Logger, authFunc grpc_auth.AuthFunc) *grpc.Server {
	myServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
			grpc_auth.StreamServerInterceptor(authFunc),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_auth.UnaryServerInterceptor(authFunc),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	return myServer
}

func WrapRequestIDAuthFunc(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, contract.RequestID, RequestIDFromIncomingContext(ctx)), nil
}
