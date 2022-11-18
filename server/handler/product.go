package handler

import "github.com/Xanvial/tutorial-grpc/server/usecase"

type ProductServer struct {
	productUC usecase.ProductClass

	// put other variable needed for grpc here
}

func NewProductHandler(productUsecase usecase.ProductClass) *ProductServer {
	return &ProductServer{
		productUC: productUsecase,
	}
}

// Put all other grpc handlers in here
