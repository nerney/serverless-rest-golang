package data

import (
	"reflect"
	"testing"

	"github.com/nerney/serverless-rest-golang/models"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		want []models.Item
	}{
		{"Should Get All",
			[]models.Item{
				{ID: "1", Data: "one"},
				{ID: "2", Data: "two"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAll(); !reflect.DeepEqual(got, tt.want) {
				if !reflect.DeepEqual(got, []models.Item{
					{
						ID:   "2",
						Data: "two",
					},
					{
						ID:   "1",
						Data: "one",
					}}) {
					t.Errorf("GetAll() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *models.Item
	}{
		{"Should Get One", args{id: "1"}, &models.Item{ID: "1", Data: "one"}},
		{"Should Not Get One", args{id: "666"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOne(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPut(t *testing.T) {
	type args struct {
		item models.Item
	}
	tests := []struct {
		name string
		args args
	}{
		{"Should Put", args{item: models.Item{ID: "666", Data: "beast"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Put(tt.args.item); !reflect.DeepEqual(got, &models.Item{ID: "666", Data: "beast"}) {
				t.Errorf("Put() wrong %v", got)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Should Delete", args{id: "666"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delete(tt.args.id)
		})
	}
}

func TestExportIfAltered(t *testing.T) {
	c.Altered = false
	if ExportIfAltered() != nil {
		t.Fail()
	}
	Put(models.Item{ID: "38942398", Data: "dataatatat"})
	if ExportIfAltered() == nil {
		t.Fail()
	}
}
