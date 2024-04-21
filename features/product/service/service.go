package service

import (
	"errors"
	"lendra/features/product"
)

type productService struct {
	productData product.ProductDataInterface
}

// dependency injection
func New(repo product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: repo,
	}
}

// Create implements product.ProductServiceInterface.
func (p *productService) Create(userIdLogin int, input product.Product) error {
	if input.NameProduct == "" {
		return errors.New("nama produk tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga produk harus lebih besar dari 0")
	}

	err := p.productData.Create(userIdLogin, input)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements product.ProductServiceInterface.
func (p *productService) Delete(userIdLogin int, idProduct int) error {
	err := p.productData.Delete(userIdLogin, idProduct)
	if err != nil {
		return err
	}

	return nil
}

// GetAllProduct implements product.ProductServiceInterface.
func (p *productService) GetAllProduct() ([]product.Product, error) {
	products, err := p.productData.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductById implements product.ProductServiceInterface.
func (p *productService) GetProductById(idProduct int) (*product.Product, error) {
	result, err := p.productData.GetProductById(idProduct)
	return result, err
}

// Update implements product.ProductServiceInterface.
func (p *productService) Update(userIdLogin, idProduct int, input product.Product) error {
	err := p.productData.Update(userIdLogin, idProduct, input)
	if err != nil {
		return err
	}
	return nil
}