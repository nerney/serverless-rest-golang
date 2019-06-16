package storage

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nerney/serverless-rest-golang/data"
	"github.com/nerney/serverless-rest-golang/models"
)

var (
	bucket = aws.String("serverless-rest-golang-items")
	key    = aws.String("items-cache")
	client = s3.New((session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))))
)

// Sync datastore with bucket
func Sync() {
	storageItems := getItemsFromStorage()
	memItems := data.ExportIfAltered()
	
	
	
	; items != nil {
		if b, err := json.Marshal(items); err == nil {
			sync(b)
		}
	}
}

func getItemsFromStorage() []models.Item {
	if o, err := client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	}); err == nil {
		defer o.Body.Close()
		storage := []byte{}
		if _, err := o.Body.Read(storage); err != nil {
			return nil
		}
		items := []models.Item{}
		if json.Unmarshal(storage, &items) != nil {
			return nil
		}
		return items
	}
	return nil
}

func sync(mem []byte) {
	storageItems := getItemsFromStorage()

		client.PutObject(&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
			Body:   mergeBytes(mem, storage),
		})
	}
}

func mergeBytes(mem, storage []byte) *bytes.Reader {

	memItems := []models.Item{}
	if err := json.Unmarshal(mem, &memItems); err != nil {
		panic(err)
	}

	storageItems := []models.Item{}
	if err := json.Unmarshal(storage, &storageItems); err != nil {
		panic(err)
	}

	merged := mergeItems(memItems, storageItems)

	b, err := json.Marshal(merged)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(b)
}

func mergeItems(memItems, storageItems []models.Item) []models.Item {
	itemMap := map[string]models.Item{}
	for _, item := range append(memItems, storageItems...) {
		itemMap[item.ID] = item
	}
	merged := []models.Item{}
	for _, item := range itemMap {
		merged = append(merged, item)
	}
	return merged
}
