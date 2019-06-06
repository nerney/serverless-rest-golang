package handlers

import (
	"context"
	"encoding/json"
	"go-sls-rest/models"
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
			if got := get(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("get() = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

func Test_post(t *testing.T) {
	body, _ := json.Marshal(models.ItemTxt{Txt: "hello"})
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
			if got := post(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
				t.Errorf("post() = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

func Test_put(t *testing.T) {
	body, _ := json.Marshal(models.ItemTxt{Txt: "hello"})
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
		{"Should fail not found",
			args{req: models.Request{
				PathParameters: map[string]string{"id": "69420"},
				Body:           string(body)}},
			models.Response{StatusCode: http.StatusBadRequest}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := put(tt.args.req); !reflect.DeepEqual(got.StatusCode, tt.want.StatusCode) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delete(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRest(t *testing.T) {
	type args struct {
		in0 context.Context
		req models.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"Should GET",
			args{in0: nil, req: models.Request{
				HTTPMethod: "get",
			}}},
		{"Should POST",
			args{in0: nil, req: models.Request{
				HTTPMethod: "post",
			}}},
		{"Should PUT",
			args{in0: nil, req: models.Request{
				HTTPMethod: "put",
			}}},
		{"Should DELETE",
			args{in0: nil, req: models.Request{
				HTTPMethod: "delete",
			}}},
		{"Should fail method not allowed",
			args{in0: nil, req: models.Request{
				HTTPMethod: "options",
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Rest(tt.args.in0, tt.args.req)
			if tt.args.req.HTTPMethod == "options" {
				if !reflect.DeepEqual(got, methodNotAllowedResponse) {
					t.Errorf("Rest() = %v, want %v", got.StatusCode, http.StatusMethodNotAllowed)
				}
			} else {
				if got.StatusCode == http.StatusMethodNotAllowed {
					t.Errorf("Rest() = %v, want %v", got.StatusCode, "Any Status Code except 405")
				}
			}
		})
	}
}
