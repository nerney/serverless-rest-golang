package storage

import (
	"bytes"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nerney/serverless-rest-golang/models"
)

var (
	bucket = aws.String("serverless-rest-golang-items")
	key    = aws.String("items-cache")
	client = s3.New((session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))))
)

// SyncToStorage merges the local cache with what is currently in s3 and then stores to s3
func SyncToStorage(items []models.Item) {
	// merge the cache from storage and the current cache
	storage := getStorage()
	if storage != nil {
		storedCache := transformStorageToCache(storage)
		fullCache := merge(storedCache, items)
		putStorage(fullCache)
	}
	putStorage(items)
}

// GetStoredItems returns stored items from s3
func GetStoredItems() []models.Item {
	b := getStorage()
	if b != nil && len(b) > 0 {
		return transformStorageToCache(b)
	}
	return nil
}

func getStorage() []byte {
	r, err := client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err == nil {
		defer r.Body.Close()
		b := new(bytes.Buffer)
		b.ReadFrom(r.Body)
		if b.Len() > 0 {
			return b.Bytes()
		}
	}
	return nil
}

func putStorage(items []models.Item) {
	body := bytes.NewReader(transformCacheToStorage(items))
	if _, err := client.PutObject(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
		Body:   body,
	}); err != nil {
		panic(err)
	}
}

func transformStorageToCache(storage []byte) []models.Item {
	items := []models.Item{}
	json.Unmarshal(storage, &items)
	return nil
}

func transformCacheToStorage(cache []models.Item) []byte {
	jsonBytes := []byte{'['}
	for _, item := range cache {
		b, _ := json.Marshal(item)
		jsonBytes = append(jsonBytes, b...)
		jsonBytes = append(jsonBytes, ',')
	}
	return append(jsonBytes, ']')
}

func merge(c1 []models.Item, c2 []models.Item) []models.Item {
	if c1 == nil && c2 == nil {
		return []models.Item{}
	}
	if c1 == nil {
		return c2
	} else if c2 == nil {
		return c1
	}

	merged := map[string]models.Item{}
	fullSlice := append(c1, c2...)
	mergedSlice := []models.Item{}
	for _, item := range fullSlice {
		merged[item.ID] = item
	}
	for _, item := range merged {
		mergedSlice = append(mergedSlice, item)
	}
	return mergedSlice
}
