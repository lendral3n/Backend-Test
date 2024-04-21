package handler

import "lendra/features/product"

type ProductRequest struct {
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
	UserId      uint
}

func RequestToCore(input ProductRequest, userIdLogin uint) product.Product {
	return product.Product{
		UserID:      userIdLogin,
		NameProduct: input.NameProduct,
		Price:       input.Price,
	}
}
