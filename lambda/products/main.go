package main

import (
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/core/product"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (product.GetProductsResponse, error) {
	logger := app.InitLogger()

	if err := app.InitConfigEnv(); err != nil {
		logger.Printf("Init config error: %+v", err)
		panic(err)
	}
	logger.Printf("Initial app config: %+v", app.CFG)

	response, _ := product.GetProducts()
	return response, nil
}

func main() {
	lambda.Start(handler)
}
