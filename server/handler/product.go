package handler

import (
	"context"

	appproto "github.com/Xanvial/tutorial-grpc/proto"
	"github.com/Xanvial/tutorial-grpc/server/model"
	"github.com/Xanvial/tutorial-grpc/server/usecase"
)

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
	appproto.UnsafeProductServiceServer
}

func NewProductHandler(productUsecase usecase.ProductClass) *ProductServer {
	return &ProductServer{
		productUC: productUsecase,
	}
}

// Put all other grpc handlers in here

func (ps *ProductServer) AddProduct(ctx context.Context, in *appproto.AddProductReq) (*appproto.AddProductResp, error) {
	err := ps.productUC.AddProduct(model.Product{
		ID:          int(in.GetProduct().GetId()),
		Name:        in.GetProduct().GetName(),
		Description: in.GetProduct().GetDescription(),
	})
	if err != nil {
		return &appproto.AddProductResp{
			Success: false,
		}, err
	}

	return &appproto.AddProductResp{
		Success: true,
	}, nil
}

func (ps *ProductServer) GetProducts(ctx context.Context, in *appproto.GetProductsReq) (*appproto.GetProductsResp, error) {
	resp := ps.productUC.GetProducts()

	ret := make([]*appproto.Product, len(resp))
	for k, v := range resp {
		ret[k] = &appproto.Product{
			Id:          int64(v.ID),
			Name:        v.Name,
			Description: v.Description,
		}
	}

	return &appproto.GetProductsResp{
		Products: ret,
	}, nil
}

func (ps *ProductServer) GetProduct(ctx context.Context, in *appproto.GetProductReq) (*appproto.GetProductResp, error) {
	resp, err := ps.productUC.GetProduct(int(in.GetId()))
	if err != nil {
		return &appproto.GetProductResp{}, err
	}

	return &appproto.GetProductResp{
		Product: &appproto.Product{
			Id:          int64(resp.ID),
			Name:        resp.Name,
			Description: resp.Description,
		},
	}, nil
}
