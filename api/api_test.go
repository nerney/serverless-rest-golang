package api

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/nerney/serverless-rest-golang/models"
)

func TestRest(t *testing.T) {
	type args struct {
		req models.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"Should GET",
			args{req: models.Request{
				HTTPMethod: "get",
			}}},
		{"Should POST",
			args{req: models.Request{
				HTTPMethod: "post",
			}}},
		{"Should PUT",
			args{req: models.Request{
				HTTPMethod: "put",
			}}},
		{"Should DELETE",
			args{req: models.Request{
				HTTPMethod: "delete",
			}}},
		{"Should fail method not allowed",
			args{req: models.Request{
				HTTPMethod: "options",
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rest(tt.args.req)
			if tt.args.req.HTTPMethod == "options" {
				if !reflect.DeepEqual(got, models.Response{StatusCode: http.StatusMethodNotAllowed}) {
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
