package product

import (
	"kiren-backend-go/internal/app"
)

// Product represents the product model
type Product struct {
	ID               string  `json:"product_id"`
	NameTH           string  `json:"product_name_th"`
	NameEN           string  `json:"product_name_en"`
	DescriptionTH    string  `json:"product_description_th"`
	DescriptionEN    string  `json:"product_description_en"`
	Price            float64 `json:"product_price"`
	ProductImagePath string  `json:"product_image_path"`
	BestSellerFlag   bool    `json:"best_seller_flag"`
	CategoryID       string  `json:"category_id"`
}

// GetProductsResponse represents the response model of GetProducts endpoint.
type GetProductsResponse struct {
	ProductData []Product     `json:"product_data"`
	Error       app.ErrorResp `json:"error,omitempty"`
}
