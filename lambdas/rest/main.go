package main

import (
	"github.com/nerney/serverless-rest-golang/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.Rest)
}
