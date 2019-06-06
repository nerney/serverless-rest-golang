package data

import (
	"github.com/nerney/serverless-rest-golang/models"
	"os"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func init() {
	c = cache.New(cache.NoExpiration, cache.NoExpiration)

	// throw a couple things in the cache for testing
	if os.Getenv("TESTING") != "" {
		c.SetDefault("1", models.Item{ID: "1", Txt: models.ItemTxt{Txt: "one"}})
		c.SetDefault("2", models.Item{ID: "2", Txt: models.ItemTxt{Txt: "two"}})
	}
}

// GetAll items from the data store.
func GetAll() []models.Item {
	items := []models.Item{}
	for _, item := range c.Items() {
		items = append(items, item.Object.(models.Item))
	}
	return items
}

// GetOne item by id from the data store.
func GetOne(id string) *models.Item {
	if item, _ := c.Get(id); item != nil {
		i := item.(models.Item)
		return &i
	}
	return nil
}

// Puts a new item into the data store or updates an existing one.
func Put(item models.Item) *models.Item {
	c.SetDefault(item.ID, item)
	return &item
}

// Delete an item from the data store by id
func Delete(id string) {
	c.Delete(id)
}
