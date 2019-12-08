package main

import (
	"bytes"
	"encoding/json"
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/core/product"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	// Initialize logger and load environment variables
	logger := app.InitLogger()
	if err := app.InitConfigEnv(); err != nil {
		logger.Errorf("Init config error: %+v", err)
		panic(err)
	}
	logger.Infof("Initial app config: %+v", app.CFG)

	// Query products data from the database
	bestSellerProducts, statusCode := product.GetBestSellerProducts()

	// Convert bestSellerProducts type from struct to string
	body, err := json.Marshal(bestSellerProducts)
	if err != nil {
		logger.Errorf("Cannot convert logic response to []byte: %+v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)

	// Return response with AWS Lambda Proxy Response
	response := events.APIGatewayProxyResponse{
		Body:       buf.String(),
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
