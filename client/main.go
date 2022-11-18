package main

import (
	"os"
)

func main() {

	// optional, create interceptor class here, and send to Product Client

	// init product client class
	productCli := NewProductClient(os.Stdin)

	productCli.MainLoop()
}
