package models

// Request type
type Request struct {
	HTTPMethod     string            `json:"httpMethod"`
	PathParameters map[string]string `json:"pathParameters,omitempty"`
	Body           string            `json:"body,omitempty"`
}

// Response type
type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body,omitempty"`
}

// Item resource
type Item struct {
	ID   string      `json:"id,omitempty"`
	Data interface{} `json:"data"`
}
