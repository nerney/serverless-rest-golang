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
	if items := data.ExportIfAltered(); items != nil {
		if b, err := json.Marshal(items); err != nil {
			panic(err)
		} else {
			bucketSync(b)
		}
	}
}

func bucketSync(mem []byte) {
	if o, err := client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	}); err != nil {
		panic(err)
	} else {
		defer o.Body.Close()
		storage := []byte{}
		if _, err := o.Body.Read(storage); err != nil {
			panic(err)
		}
		client.PutObject(&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
			Body:   merge(mem, storage),
		})
	}
}

func merge(mem, storage []byte) *bytes.Reader {
	merged := []models.Item{}

	items := []models.Item{}
	if err := json.Unmarshal(mem, &items); err != nil {
		panic(err)
	}

	merged = append(merged, items...)

	items = []models.Item{}
	if err := json.Unmarshal(storage, &items); err != nil {
		panic(err)
	}

	merged = append(merged, items...)

	itemMap := map[string]models.Item{}
	for _, item := range merged {
		itemMap[item.ID] = item
	}

	merged = []models.Item{}
	for _, item := range itemMap {
		merged = append(merged, item)
	}

	b, err := json.Marshal(merged)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(b)
}
