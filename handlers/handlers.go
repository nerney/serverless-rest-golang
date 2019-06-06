package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-sls-rest/models"
	"go-sls-rest/service"
	"net/http"
	"strings"
)

var (
	methodNotAllowedResponse = models.Response{StatusCode: http.StatusMethodNotAllowed}
	notFoundResponse         = models.Response{StatusCode: http.StatusNotFound}
	badRequestResponse       = models.Response{StatusCode: http.StatusBadRequest}
)

// Builds an ok response with a body.
func okResponse(body interface{}) models.Response {
	return models.Response{StatusCode: http.StatusOK, Body: fmt.Sprintf("%v", body)}
}

// Handles get one and get all requests
func get(req models.Request) models.Response {
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

// Handles post requests
func post(req models.Request) models.Response {
	var txt models.ItemTxt
	if json.Unmarshal([]byte(req.Body), &txt) != nil {
		return badRequestResponse
	}
	body, err := json.Marshal(service.Create(txt))
	if err != nil {
		panic(err)
	}
	return okResponse(string(body))
}

// Handles put requests
func put(req models.Request) models.Response {
	id := req.PathParameters["id"]
	if id == "" {
		return badRequestResponse
	}
	var txt models.ItemTxt
	if json.Unmarshal([]byte(req.Body), &txt) != nil {
		return badRequestResponse
	}
	item := models.Item{ID: id, Txt: txt}
	if service.Update(item) != nil {
		return badRequestResponse
	}
	body, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return okResponse(string(body))
}

// Handles delete requests
func delete(req models.Request) models.Response {
	id := req.PathParameters["id"]
	if id == "" {
		return badRequestResponse
	}
	if service.Delete(id) != nil {
		return notFoundResponse
	}
	return models.Response{StatusCode: http.StatusNoContent}
}

// Rest handler will accept an incoming request
// and pass it along to the appropriate method handler
// which will then return a response.
func Rest(_ context.Context, req models.Request) (models.Response, error) {
	switch strings.ToUpper(req.HTTPMethod) {
	case "GET":
		return get(req), nil
	case "POST":
		return post(req), nil
	case "PUT":
		return put(req), nil
	case "DELETE":
		return delete(req), nil
	}
	return methodNotAllowedResponse, nil
}
