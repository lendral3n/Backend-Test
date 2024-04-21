package data

import (
	"lendra/features/product"
	"lendra/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NameProduct string
	Price       int
	UserID      uint
	User        data.User
}

func ProductToModel(input product.Product) Product {
	return Product{
		NameProduct: input.NameProduct,
		Price:       input.Price,
	}
}

func (u Product) ModelToProduct() product.Product {
	return product.Product{
		ID:          u.ID,
		NameProduct: u.NameProduct,
		Price:       u.Price,
		UserID:      u.UserID,
	}
}
