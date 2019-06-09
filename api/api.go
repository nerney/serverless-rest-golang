package api

import (
	"net/http"
	"strings"

	"github.com/nerney/serverless-rest-golang/handlers"
	"github.com/nerney/serverless-rest-golang/models"
)

// Rest api will accept an incoming request
// and pass it along to the appropriate method handler
// which will then return a response.
func Rest(req models.Request) models.Response {
	switch strings.ToUpper(req.HTTPMethod) {
	case "GET":
		return handlers.Get(req)
	case "POST":
		return handlers.Post(req)
	case "PUT":
		return handlers.Put(req)
	case "DELETE":
		return handlers.Delete(req)
	}
	return models.Response{StatusCode: http.StatusMethodNotAllowed}
}
