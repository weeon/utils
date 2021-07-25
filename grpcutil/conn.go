package grpcutil

import (
	"context"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
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

func NewServer(logger *zap.Logger, authFunc grpc_auth.AuthFunc, opt ...grpc.ServerOption) *grpc.Server {
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_opentracing.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(logger),
		grpc_auth.StreamServerInterceptor(authFunc),
		grpc_recovery.StreamServerInterceptor(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_auth.UnaryServerInterceptor(authFunc),
			grpc_recovery.UnaryServerInterceptor(),
			UnaryServerWrapRequestIDInterceptor(),
		)))
	if opt != nil {
		opts = append(opts, opt...)
	}

	myServer := grpc.NewServer(
		opts...,
	)
	return myServer
}

func WrapRequestIDAuthFunc(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, contract.RequestID, RequestIDFromIncomingContext(ctx)), nil
}
