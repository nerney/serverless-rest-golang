package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nerney/serverless-rest-golang/api"
	"github.com/nerney/serverless-rest-golang/data"
	"github.com/nerney/serverless-rest-golang/models"
	"github.com/nerney/serverless-rest-golang/storage"
)

// where the lambda enters
func main() {
	lambda.Start(lambdaWrapper)
}

// LambdaWrapper abstracts the REST application logic away from the signature required by AWS API Gateway / Lambda
func lambdaWrapper(_ context.Context, req models.Request) (models.Response, error) {
	defer storage.SyncToStorage(data.GetAll())
	return api.Rest(req), nil
}
