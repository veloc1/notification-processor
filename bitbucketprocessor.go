package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BitbucketLink struct {
	Href string
}

type BitbucketLinks struct {
	Comments BitbucketLink
	Html     BitbucketLink
}

type BitbucketPR struct {
	Links BitbucketLinks
}

type BitbucketRepo struct {
	Name string
}

type BitbucketEvent struct {
	Pullrequest BitbucketPR
	Repository  BitbucketRepo
}

type BitbucketProcessor struct {
	Processor
}

func (BitbucketProcessor) canHandle(script string) bool {
	return script == "bitbucket"
}

func (BitbucketProcessor) process(r *http.Request) (NotifyData, error) {
	if r.Method != http.MethodPost {
		return NotifyData{}, errors.New("Wrong request method")
	}

	eventHeader := r.Header.Get("X-Event-Key")
	if eventHeader == "" {
		return NotifyData{}, errors.New("Wrong event key")
	}

	var decoder = json.NewDecoder(r.Body)
	var event BitbucketEvent
	err := decoder.Decode(&event)
	if err != nil {
		return NotifyData{}, errors.New("Wrong event JSON object")
	}

	project := event.Repository.Name
	message := "Pullrequest created at " + project
	message = message + "\n\n"
	message = message + "Url: <a href=\"" + event.Pullrequest.Links.Html.Href + "\">Link</a>"
	fmt.Print(message)

	return NotifyData{
		message: message,
		project: project,
		groups:  []string{"developers"},
	}, nil
}
