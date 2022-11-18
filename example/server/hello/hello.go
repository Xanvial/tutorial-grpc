package hello

import (
	"context"
	"log"

	tutorialproto "github.com/Xanvial/tutorial-grpc/example/proto"
)

type Server struct {
	// this will allow forward compatibility
	// but new function will always return error unless implemented in code
	tutorialproto.UnimplementedChatServiceServer
}

func (s *Server) TestHello(ctx context.Context, in *tutorialproto.MessageReq) (*tutorialproto.MessageResp, error) {
	log.Printf("Receive message: %s", in.GetBody())
	return &tutorialproto.MessageResp{Body: "Hello From the Server!"}, nil
}
