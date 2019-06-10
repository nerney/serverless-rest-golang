package main

import (
	"bytes"
	"fmt"

	"net/http"

	"github.com/nerney/serverless-rest-golang/api"
	"github.com/nerney/serverless-rest-golang/models"
)

// run the endpoint locally
func main() {
	fmt.Println("Server up @ :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := new(bytes.Buffer)
		b.ReadFrom(r.Body)
		res := api.Rest(models.Request{
			HTTPMethod:     r.Method,
			PathParameters: map[string]string{"id": r.URL.Path[1:]},
			Body:           b.String(),
		})
		fmt.Printf("[%v] %v --> %v\n", r.Method, r.URL.RequestURI(), res.StatusCode)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Body))
	})
	panic(http.ListenAndServe(":8080", nil))
}
