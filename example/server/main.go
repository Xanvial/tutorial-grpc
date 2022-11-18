package main

import (
	"log"
	"net"

	tutorialproto "github.com/Xanvial/tutorial-grpc/example/proto"
	"github.com/Xanvial/tutorial-grpc/example/server/hello"
	"github.com/Xanvial/tutorial-grpc/example/server/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := hello.Server{}
	grpcInterceptor := interceptor.NewGRPCInterceptor()

	// grpc unary interceptors, will be executed from first element to the last
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpcInterceptor.LoggingInterceptor(),
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)

	tutorialproto.RegisterChatServiceServer(grpcServer, &s)

	// register server using reflection
	reflection.Register(grpcServer)

	log.Println("start listening on port: 8999")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
