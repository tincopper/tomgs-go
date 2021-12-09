package rpc

import (
	"context"
	"tomgs-go/learning-grpc/pkg/errors"

	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return errors.WrapRPCError(err)
}