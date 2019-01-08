package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
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

	bbUsername string
	bbPassword string
}

func (BitbucketProcessor) canHandle(script string) bool {
	return script == "bitbucket"
}

func (processor BitbucketProcessor) process(r *http.Request) (NotifyData, error) {
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

	addPrComment(event.Pullrequest.Links.Comments.Href, processor.bbUsername, processor.bbPassword)

	return NotifyData{
		message: message,
		project: project,
		groups:  []string{"developers"},
	}, nil
}

func addPrComment(commentsUrl string, username string, password string) bool {
	client := &http.Client{}

	body := `{
		"content": {
			"raw": "sample text"
		}
	}`

	req, err := http.NewRequest("POST", commentsUrl, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(username, password)

	if err != nil {
		return false
	}
	resp, err := client.Do(req)

	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
