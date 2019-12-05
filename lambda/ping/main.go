package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func handler() (Response, error) {
	response := Response{
		Message: "pong",
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
