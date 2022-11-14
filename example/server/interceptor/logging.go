package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type GRPCInterceptor struct {
}

func NewGRPCInterceptor() *GRPCInterceptor {
	return &GRPCInterceptor{}
}

func (gi *GRPCInterceptor) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("LoggingInterceptor Request:", req)

		return handler(ctx, req)
	}
}
