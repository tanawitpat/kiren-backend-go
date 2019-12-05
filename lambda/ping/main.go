package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	response := Response{
		Message: "pong",
	}
	return response, nil
}

func main() {
	lambda.Start(Handler)
}
