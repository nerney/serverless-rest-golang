package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nerney/serverless-rest-golang/api"
	"github.com/nerney/serverless-rest-golang/models"
	"github.com/nerney/serverless-rest-golang/storage"
)

// where the lambda enters
func main() {
	defer storage.Sync()
	lambda.Start(func(_ context.Context, req models.Request) (models.Response, error) {
		return api.Rest(req), nil
	})
}
