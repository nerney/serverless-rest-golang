package data

import (
	"github.com/nerney/serverless-rest-golang/models"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		want []models.Item
	}{
		{"Should Get All",
			[]models.Item{
				{
					ID:  "1",
					Txt: models.ItemTxt{Txt: "one"},
				},
				{
					ID:  "2",
					Txt: models.ItemTxt{Txt: "two"},
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAll(); !reflect.DeepEqual(got, tt.want) {
				if !reflect.DeepEqual(got, []models.Item{
					{
						ID:  "2",
						Txt: models.ItemTxt{Txt: "two"},
					},
					{
						ID:  "1",
						Txt: models.ItemTxt{Txt: "one"},
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
		{"Should Get One", args{id: "1"}, &models.Item{ID: "1", Txt: models.ItemTxt{Txt: "one"}}},
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
	i := models.Item{ID: "666", Txt: models.ItemTxt{Txt: "beast"}}
	type args struct {
		item models.Item
	}
	tests := []struct {
		name string
		args args
		want *models.Item
	}{
		{"Should Put", args{item: i}, &i},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Put(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
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
