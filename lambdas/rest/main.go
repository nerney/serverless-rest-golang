package main

import (
	"go-sls-rest/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.Rest)
}
