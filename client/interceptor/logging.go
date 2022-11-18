package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type GRPCInterceptor struct {
}

func NewGRPCInterceptor() *GRPCInterceptor {
	return &GRPCInterceptor{}
}

func (gi *GRPCInterceptor) LoggingInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Println("process takes:", time.Since(start).String())

		return err
	}
}
