package data

import (
	"os"

	"github.com/nerney/serverless-rest-golang/models"
	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func init() {
	c = cache.New(cache.NoExpiration, cache.NoExpiration)

	// throw a couple things in a cache for testing
	if os.Getenv("TESTING") != "" {
		c.SetDefault("1", models.Item{ID: "1", Data: "one"})
		c.SetDefault("2", models.Item{ID: "2", Data: "two"})
		return
	}
}

// GetAll items from the data cache.
func GetAll() []models.Item {
	items := []models.Item{}
	for _, item := range c.Items() {
		items = append(items, item.Object.(models.Item))
	}
	return items
}

// GetOne item by id from the data cache.
func GetOne(id string) *models.Item {
	if item, _ := c.Get(id); item != nil {
		i := item.(models.Item)
		return &i
	}
	return nil
}

// Put a new item into the data cache or updates an existing one.
func Put(item models.Item) *models.Item {
	c.SetDefault(item.ID, item)
	return &item
}

// Delete an item from the data cache by id
func Delete(id string) {
	c.Delete(id)
}
