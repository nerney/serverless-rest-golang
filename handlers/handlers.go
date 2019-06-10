package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nerney/serverless-rest-golang/models"
	"github.com/nerney/serverless-rest-golang/service"
)

var (
	notFoundResponse   = models.Response{StatusCode: http.StatusNotFound}
	badRequestResponse = models.Response{StatusCode: http.StatusBadRequest}
)

// Builds an ok response with a body.
func okResponse(body interface{}) models.Response {
	return models.Response{StatusCode: http.StatusOK, Body: fmt.Sprintf("%v", body)}
}

// Get handles get one and get all requests
func Get(req models.Request) models.Response {
	var (
		body []byte
		err  error
	)
	// get one
	if id := req.PathParameters["id"]; id != "" {
		if item, err := service.GetOne(id); err == nil {
			if body, err = json.Marshal(item); err != nil {
				panic(err)
			}
			return okResponse(string(body))
		}
		return notFoundResponse
	}
	// get all
	if body, err = json.Marshal(service.GetAll()); err != nil {
		panic(err)
	}
	return okResponse(string(body))
}

// Post handles post requests
func Post(req models.Request) models.Response {
	item := models.Item{}
	if json.Unmarshal([]byte(req.Body), &item) != nil {
		return badRequestResponse
	}
	body, err := json.Marshal(service.Create(item))
	if err != nil {
		panic(err)
	}
	return okResponse(string(body))
}

// Put handles put requests
func Put(req models.Request) models.Response {
	id := req.PathParameters["id"]
	if id == "" {
		return badRequestResponse
	}
	var item models.Item
	if json.Unmarshal([]byte(req.Body), &item) != nil {
		return badRequestResponse
	}
	item = models.Item{ID: id, Data: item.Data}
	if service.Update(item) != nil {
		return badRequestResponse
	}
	body, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return okResponse(string(body))
}

// Delete handles delete requests
func Delete(req models.Request) models.Response {
	id := req.PathParameters["id"]
	if id == "" {
		return badRequestResponse
	}
	if service.Delete(id) != nil {
		return notFoundResponse
	}
	return models.Response{StatusCode: http.StatusNoContent}
}
