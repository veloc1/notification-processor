package main

import (
	"fmt"
	"net/http"
)

type WebhookHandler struct {
	http.Handler

	bitbucket Processor
}

func (h WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var values map[string][]string = r.URL.Query()
	if len(values) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var script = values["service"][0]
		switch script {
		case "bitbucket":
			h.process(h.bitbucket, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (h WebhookHandler) process(p Processor, w http.ResponseWriter) {
	_, err := p.process(1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
