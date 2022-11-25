package main

import (
	"os"

	product "github.com/Xanvial/tutorial-grpc/client/Product"
)

func main() {

	// optional, create interceptor class here, and send to Product Client

	// init product client class
	productCli := product.NewProductClient(os.Stdin)

	productCli.MainLoop()
}
