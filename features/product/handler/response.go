package handler

import "lendra/features/product"

type ProductResponse struct {
	ID          int    `json:"id" form:"id"`
	NameProduct string `json:"name_product" form:"name_product"`
	Price       int    `json:"price" form:"price"`
	UserID      uint   `json:"user_id"`
}

func CoreToResponse(data product.Product) ProductResponse {
	var result = ProductResponse{
		ID:          int(data.ID),
		NameProduct: data.NameProduct,
		Price:       data.Price,
		UserID:      data.UserID,
	}
	return result
}
