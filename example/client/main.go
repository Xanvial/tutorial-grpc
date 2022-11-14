package main

import (
	"context"
	"log"

	tutorialproto "github.com/Xanvial/tutorial-grpc/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8999",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // to allow insecure connection
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := tutorialproto.NewChatServiceClient(conn)

	response, err := c.TestHello(context.Background(), &tutorialproto.MessageReq{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
