package utils

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	s := fmt.Sprintf("%s : %s", info.FullMethod, req)
	fmt.Println("zz", s)
	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}
