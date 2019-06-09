package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nerney/serverless-rest-golang/api"
	"github.com/nerney/serverless-rest-golang/models"
)

// where the lambda enters
func main() {
	lambda.Start(func(_ context.Context, req models.Request) (models.Response, error) {
		//defer storage.SyncToStorage(data.GetAll())
		return api.Rest(req), nil
	})
}
