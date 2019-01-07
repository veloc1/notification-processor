package main

import (
	"net/http"
)

func main() {
	handler := &WebhookHandler{
		processors: []Processor{
			&BitbucketProcessor{},
		},
	}
	http.Handle("/", handler)

	http.ListenAndServe(":9180", nil)
}
