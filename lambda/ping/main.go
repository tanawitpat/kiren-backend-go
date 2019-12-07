package main

import (
	"kiren-backend-go/internal/core/pingpong"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (events.APIGatewayProxyResponse, error) {
	body, statusCode := pingpong.Ping()
	response := events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
