package main

import (
	"kiren-backend-go/internal/core/pingpong"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
	response, _ := pingpong.Ping()
	return response, nil
}

func main() {
	lambda.Start(handler)
}
