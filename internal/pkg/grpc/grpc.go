package grpc

import (
	"context"
	"crypto/tls"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	utils_context "tinvest-go/internal/pkg/context"
)

func NewConnection(appName string, target string, timeout time.Duration) grpc.ClientConnInterface {
	params := strings.Split(target, ":")

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: params[0],
			MinVersion: tls.VersionTLS13,
		})),
		grpc.WithUserAgent(appName),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{PermitWithoutStream: true}),
		grpc.WithChainUnaryInterceptor(
			NewTimeoutUnaryInterceptor(timeout),
			NewAppNameUnaryInterceptor(appName),
		),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		panic(err)
	}

	return conn
}

func NewAppNameUnaryInterceptor(appName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := metadata.AppendToOutgoingContext(ctx, utils_context.AppNameHeader, appName)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func NewTimeoutUnaryInterceptor(t time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, t)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
