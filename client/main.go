package main

import (
	"os"

	product "github.com/Xanvial/tutorial-grpc/client/Product"
	"github.com/Xanvial/tutorial-grpc/client/interceptor"
)

func main() {

	// optional, create interceptor class here, and send to Product Client
	grpcInterceptor := interceptor.NewGRPCInterceptor()

	// init product client class
	productCli := product.NewProductClient(os.Stdin, grpcInterceptor)

	productCli.MainLoop()
}
