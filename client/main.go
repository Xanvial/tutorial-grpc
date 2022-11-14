package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // to allow insecure connection
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// put logic here
}
