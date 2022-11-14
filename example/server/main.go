package main

import (
	"log"
	"net"

	tutorialproto "github.com/Xanvial/tutorial-grpc/example/proto"
	"github.com/Xanvial/tutorial-grpc/example/server/hello"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := hello.Server{}

	tutorialproto.RegisterChatServiceServer(grpcServer, &s)

	log.Println("start listening on port: 8999")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
