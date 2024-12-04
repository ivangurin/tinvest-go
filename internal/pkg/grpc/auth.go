package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type Auth struct {
	Token string
}

func NewAuth(token string) grpc.CallOption {
	return grpc.PerRPCCredentials(
		Auth{
			Token: token,
		})
}

func (a Auth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + a.Token,
	}, nil
}

func (a Auth) RequireTransportSecurity() bool {
	return true
}
