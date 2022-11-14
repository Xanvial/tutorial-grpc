package usecase

type ProductUsecase struct {
}

// NewProductHandler will create a new class for product usecase
func NewProductHandler() ProductUsecase {
	return ProductUsecase{}
}

// GetProduct will return list of product as specified
func (pu *ProductUsecase) GetProduct() {

}
