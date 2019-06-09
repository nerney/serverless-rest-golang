package service

import (
	"errors"

	"github.com/nerney/serverless-rest-golang/data"
	"github.com/nerney/serverless-rest-golang/models"

	"github.com/teris-io/shortid"
)

// GetAll items or return an empty array.
func GetAll() []models.Item {
	return data.GetAll()
}

// GetOne item by id, returns an error if not found.
func GetOne(id string) (models.Item, error) {
	if item := data.GetOne(id); item != nil {
		return *item, nil
	}
	return models.Item{}, errors.New("not found")
}

// Create a new item and return it with the newly generated id.
func Create(item models.Item) models.Item {
	return *data.Put(models.Item{ID: shortid.MustGenerate(), Data: item})
}

// Update an existing item, returns an error if not found.
func Update(item models.Item) error {
	if _, err := GetOne(item.ID); err != nil {
		return err
	}
	data.Put(item)
	return nil
}

// Delete an item, returns an error if not found.
func Delete(id string) error {
	if _, err := GetOne(id); err != nil {
		return err
	}
	data.Delete(id)
	return nil
}
