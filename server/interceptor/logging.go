package interceptor

import (
	"context"
	"log"

	appproto "github.com/Xanvial/tutorial-grpc/proto"
	"google.golang.org/grpc"
)

type GRPCInterceptor struct {
}

func NewGRPCInterceptor() *GRPCInterceptor {
	return &GRPCInterceptor{}
}

func (gi *GRPCInterceptor) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		switch v := req.(type) {
		case *appproto.AddProductReq:
			log.Println("LoggingInterceptor Add product request:", v)
		case *appproto.GetProductsReq:
			log.Println("LoggingInterceptor Get all products request")
		case *appproto.GetProductReq:
			log.Println("LoggingInterceptor Get product request by id:", v.GetId())
		}

		return handler(ctx, req)
	}
}
