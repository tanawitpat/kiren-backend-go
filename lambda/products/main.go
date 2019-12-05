package main

import (
	"database/sql"
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/core/product"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (product.Products, error) {
	logger := app.InitLogger()
	if err := app.InitConfigEnv(); err != nil {
		logger.Printf("Init config error: %+v", err)
		panic(err)
	}
	logger.Printf("Initial app config: %+v", app.CFG)

	products := product.Products{}
	selDB, err := app.Query(`
		SELECT product_id, product_name_th, product_description_th, product_image_path, product_price 
		FROM products 
	`)
	for selDB.Next() {
		var productID, productNameTH, productNameEN, productDescriptionTH, productDescriptionEN, productImagePath sql.NullString
		var productPrice sql.NullFloat64
		var productBestSellerFlag sql.NullBool
		err = selDB.Scan(&productID, &productNameTH, &productDescriptionTH, &productImagePath, &productPrice)
		if err != nil {
			return products, err
		}
		productSQLNull := product.ProductSQLNull{
			ID:               productID,
			NameTH:           productNameTH,
			NameEN:           productNameEN,
			DescriptionTH:    productDescriptionTH,
			DescriptionEN:    productDescriptionEN,
			ProductImagePath: productImagePath,
			Price:            productPrice,
			BestSellerFlag:   productBestSellerFlag,
		}
		product := product.MapSQLNullToProduct(productSQLNull)
		products = append(products, product)
	}
	return products, nil
}

func main() {
	lambda.Start(handler)
}
