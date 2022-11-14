package hello

import (
	"context"
	"log"

	tutorialproto "github.com/Xanvial/tutorial-grpc/example/proto"
)

type Server struct {
}

func (s *Server) TestHello(ctx context.Context, in *tutorialproto.MessageReq) (*tutorialproto.MessageResp, error) {
	log.Printf("Receive message: %s", in.Body)
	return &tutorialproto.MessageResp{Body: "Hello From the Server!"}, nil
}
