package main

import (
	"fmt"
	"net/http"
)

type WebhookHandler struct {
	http.Handler

	processors []Processor
}

func (h WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var values map[string][]string = r.URL.Query()
	if len(values) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var script = values["service"][0]
		isProcessed := false
		for _, processor := range h.processors {
			if processor.canHandle(script) {
				h.process(processor, w, r)
				isProcessed = true
			}
		}
		if !isProcessed {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (h WebhookHandler) process(p Processor, w http.ResponseWriter, r *http.Request) {
	_, err := p.process(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
