package usecase

import "github.com/Xanvial/tutorial-grpc/server/model"

type ProductClass interface {
	AddProduct(product model.Product) error
	UpdateProduct(product model.Product) error
	GetProducts() []model.Product
	GetProduct(id int) (*model.Product, error)
}
