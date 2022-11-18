package handler

import "github.com/Xanvial/tutorial-grpc/server/usecase"

type ProductServer struct {
	productUC usecase.ProductClass

	// UnsafeProductServiceServer is used for grpc forward-compatibility.
	//
	// in the case of adding new method in proto file, there's two approach for forward-compatibility:
	//
	// 1. like in example/server/hello/hello.go, adding UnimplementedXXXX will always make the compile successfull
	// even if there's no implementation for new method.
	// But on runtime, if the new method is called, it will automatically return error.
	//
	// 2. this approach, adding UnsafeXXXX will force compile error if new method is not implemented

	// uncomment this, after correctly import the proto file
	// appproto.UnsafeProductServiceServer
}

func NewProductHandler(productUsecase usecase.ProductClass) *ProductServer {
	return &ProductServer{
		productUC: productUsecase,
	}
}

// Put all other grpc handlers in here
