package service

import (
	"go-sls-rest/models"

	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		want []models.Item
	}{
		{"Should Get All", []models.Item{
			models.Item{
				ID:  "1",
				Txt: models.ItemTxt{Txt: "one"},
			},
			models.Item{
				ID:  "2",
				Txt: models.ItemTxt{Txt: "two"},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAll(); !reflect.DeepEqual(got, tt.want) {
				if !reflect.DeepEqual(got, []models.Item{
					models.Item{
						ID:  "2",
						Txt: models.ItemTxt{Txt: "two"},
					},
					models.Item{
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
		name    string
		args    args
		want    models.Item
		wantErr bool
	}{
		{"Should Get One", args{id: "1"}, models.Item{ID: "1", Txt: models.ItemTxt{Txt: "one"}}, false},
		{"Should Not Get One", args{id: "999"}, models.Item{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	txt := models.ItemTxt{Txt: "test"}
	type args struct {
		txt models.ItemTxt
	}
	tests := []struct {
		name string
		args args
		want models.ItemTxt
	}{
		{"Should Create", args{txt: txt}, txt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.txt); !reflect.DeepEqual(got.Txt, tt.want) {
				t.Errorf("Create() = %v, wanted with text %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		item models.Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Should Update", args{item: models.Item{ID: "2"}}, false},
		{"Should Not Update", args{item: models.Item{ID: "69"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Should Delete", args{id: "2"}, false},
		{"Should Not Delete", args{id: "420"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
