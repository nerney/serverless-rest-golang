package models

import (
	"github.com/aws/aws-lambda-go/events"
)

// Request from API Gateway
type Request events.APIGatewayProxyRequest

// Response from API Gateway
type Response events.APIGatewayProxyResponse

// Item resource
type Item struct {
	ID  string  `json:"id"`
	Txt ItemTxt `json:"txt"`
}

// ItemTxt content for Item struct
type ItemTxt struct {
	Txt string `json:"txt"`
}
