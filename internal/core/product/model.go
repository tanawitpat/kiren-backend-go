package product

import (
	"database/sql"
	"kiren-backend-go/internal/app"
)

// Products represents the products model.
type Products []Product

// Product represents the product model.
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

// ProductSQLNull represents the product model with SQL Null schema
type ProductSQLNull struct {
	ID               sql.NullString  `json:"product_id"`
	NameTH           sql.NullString  `json:"product_name_th"`
	NameEN           sql.NullString  `json:"product_name_en"`
	DescriptionTH    sql.NullString  `json:"product_description_th"`
	DescriptionEN    sql.NullString  `json:"product_description_en"`
	Price            sql.NullFloat64 `json:"product_price"`
	ProductImagePath sql.NullString  `json:"product_image_path"`
	BestSellerFlag   sql.NullBool    `json:"best_seller_flag"`
	CategoryID       sql.NullString  `json:"category_id"`
}

// GetProductsResponse represents the response model of GetProducts endpoint.
type GetProductsResponse struct {
	ProductData []Product     `json:"product_data,omitempty"`
	Error       app.ErrorResp `json:"error,omitempty"`
}

// GetProductResponse represents the response model of GetProduct endpoint.
type GetProductResponse struct {
	ProductData         Product       `json:"product_data,omitempty"`
	RelevantProductData []Product     `json:"relevant_product_data,omitempty"`
	Error               app.ErrorResp `json:"error,omitempty"`
}

// BestSellerProductResponse represents the response model of GetBestSellerProduct endpoint.
type BestSellerProductResponse struct {
	ProductData []Product     `json:"product_data,omitempty"`
	Error       app.ErrorResp `json:"error,omitempty"`
}
