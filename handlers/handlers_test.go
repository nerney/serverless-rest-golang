package handlers

import (
	"encoding/json"
	"github.com/nerney/serverless-rest-golang/models"
	"net/http"
	"reflect"
	"testing"
)

func Test_okResponse(t *testing.T) {
	type args struct {
		body interface{}
	}
	tests := []struct {
		name string
		args args
		want models.Response
	}{
		{"Build OK Response",
			args{body: "test"},
			models.Response{StatusCode: http.StatusOK, Body: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := okResponse(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("okResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	type args struct {
		req models.Request
	}
	tests := []struct {
		name string
		args args
		want models.Response
	}{
		{"Should get one",
			args{req: models.Request{
				PathParameters: map[string]string{"id": "1"}}},
			models.Response{StatusCode: http.StatusOK}},
		{"Should get all",
			args{req: models.Request{PathParameters: nil}},
			models.Response{StatusCode: http.StatusOK}},
		{"Should fail not found",
			args{req: models.Request{PathParameters: map[string]string{"id": "420"}}},
			models.Response{StatusCode: http.StatusNotFound}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("get() = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

func Test_post(t *testing.T) {
	body, _ := json.Marshal(models.Item{Data: "hello"})
	type args struct {
		req models.Request
	}
	tests := []struct {
		name string
		args args
		want models.Response
	}{
		{"Should post item",
			args{req: models.Request{Body: string(body)}},
			models.Response{StatusCode: http.StatusOK}},
		{"Should fail bad request",
			args{req: models.Request{Body: "fail"}},
			models.Response{StatusCode: http.StatusBadRequest}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Post(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("post() = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

func Test_put(t *testing.T) {
	body, _ := json.Marshal(models.Item{Data: "hello"})
	type args struct {
		req models.Request
	}
	tests := []struct {
		name string
		args args
		want models.Response
	}{
		{"Should put item",
			args{req: models.Request{
				PathParameters: map[string]string{"id": "2"},
				Body:           string(body)}},
			models.Response{StatusCode: http.StatusOK}},
		{"Should fail bad request",
			args{req: models.Request{
				PathParameters: map[string]string{"id": "2"},
				Body:           "fail"}},
			models.Response{StatusCode: http.StatusBadRequest}},
		{"Should fail bad request",
			args{req: models.Request{
				PathParameters: nil,
				Body:           "fail"}},
			models.Response{StatusCode: http.StatusBadRequest}},
		{"Should fail not found",
			args{req: models.Request{
				PathParameters: map[string]string{"id": "69420"},
				Body:           string(body)}},
			models.Response{StatusCode: http.StatusBadRequest}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Put(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("put() = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

func Test_delete(t *testing.T) {
	type args struct {
		req models.Request
	}
	tests := []struct {
		name string
		args args
		want models.Response
	}{
		{"Should delete item",
			args{req: models.Request{PathParameters: map[string]string{"id": "2"}}},
			models.Response{StatusCode: http.StatusNoContent}},
		{"Should fail not found",
			args{req: models.Request{PathParameters: map[string]string{"id": "999"}}},
			models.Response{StatusCode: http.StatusNotFound}},
		{"Should fail bad request",
			args{req: models.Request{PathParameters: nil}},
			models.Response{StatusCode: http.StatusBadRequest}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
