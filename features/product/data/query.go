package data

import (
	"errors"
	"lendra/features/product"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

// Create implements product.ProductDataInterface.
func (p *productQuery) Create(userIdLogin int, input product.Product) error {
	productInput := ProductToModel(input)
	productInput.UserID = uint(userIdLogin)

	tx := p.db.Create(&productInput)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements product.ProductDataInterface.
func (p *productQuery) Delete(userIdLogin int, idProduct int) error {
	var product Product
	tx := p.db.First(&product, idProduct)
	if tx.Error != nil {
		return tx.Error
	}
	if product.UserID != uint(userIdLogin) {
		return errors.New("user not authorized")
	}
	tx = p.db.Delete(&product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAllProduct implements product.ProductDataInterface.
func (p *productQuery) GetAllProduct() ([]product.Product, error) {
	var products []Product
	tx := p.db.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var result []product.Product
	for _, product := range products {
		result = append(result, product.ModelToProduct())
	}
	return result, nil
}

// GetProductById implements product.ProductDataInterface.
func (p *productQuery) GetProductById(idProduct int) (*product.Product, error) {
	var product Product
	tx := p.db.First(&product, idProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := product.ModelToProduct()
	return &result, nil
}

// Update implements product.ProductDataInterface.
func (p *productQuery) Update(userIdLogin, idProduct int, input product.Product) error {
	var product Product
	tx := p.db.First(&product, idProduct)
	if tx.Error != nil {
		return tx.Error
	}
	if product.UserID != uint(userIdLogin) {
		return errors.New("user not authorized")
	}
	product.NameProduct = input.NameProduct
	product.Price = input.Price
	tx = p.db.Save(&product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
