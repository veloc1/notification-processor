package main

import "net/http"

type Processor interface {
	canHandle(string) bool
	process(*http.Request) (NotifyData, error)
}
