package main

import (
	"kiren-backend-go/internal/core/pingpong"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (events.APIGatewayProxyResponse, error) {
	// Get the response from the pingpong.Ping function
	body, statusCode := pingpong.Ping()

	// Return response with AWS Lambda Proxy Response
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
