package usecase

import (
	"errors"

	"github.com/Xanvial/tutorial-grpc/server/model"
)

type ProductUsecase struct {
	products map[int]model.Product
}

// NewProductUsecase will create a new class for product usecase
func NewProductUsecase() *ProductUsecase {
	return &ProductUsecase{
		products: make(map[int]model.Product),
	}
}

// AddProduct will add product data to list
func (pu *ProductUsecase) AddProduct(product model.Product) error {
	// check existence to avoid duplicate id
	_, ok := pu.products[product.ID]
	if ok {
		// product already exist
		return errors.New("product id already exist")
	}

	pu.products[product.ID] = product
	return nil
}

// UpdateProduct will update existing product data
func (pu *ProductUsecase) UpdateProduct(product model.Product) error {
	// check existence
	_, ok := pu.products[product.ID]
	if !ok {
		// product haven't yet added, return error because it's supposed to use `AddProduct` for new data
		return errors.New("product id not exist")
	}

	pu.products[product.ID] = product
	return nil
}

// GetProducts will return list of product as specified
func (pu *ProductUsecase) GetProducts() []model.Product {
	// need to convert data from map to slice
	productSlice := make([]model.Product, len(pu.products))

	idx := 0
	for _, v := range pu.products {
		productSlice[idx] = v
		idx++
	}

	return productSlice
}

// GetProduct will return a product as specified
func (pu *ProductUsecase) GetProduct(id int) (*model.Product, error) {
	// check existence
	data, ok := pu.products[id]
	if !ok {
		return nil, errors.New("product id not exist")
	}

	return &data, nil
}
