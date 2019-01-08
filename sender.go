package main

import (
	"net/http"
	"net/url"
	"strings"
)

type SendMessageBody struct {
}

type Sender struct {
	manager Manager
	token   string
}

func (s Sender) Send(data NotifyData) bool {
	receivers := s.manager.GetReceivers(data)

	isSuccessful := true

	for _, receiver := range receivers {
		requestData := url.Values{}
		requestData.Add("chat_id", receiver)
		requestData.Add("text", data.message)
		requestData.Add("parse_mode", "html")

		isSended := s.request("sendMessage", requestData)
		if !isSended {
			isSuccessful = false
		}
	}

	return isSuccessful
}

func (s Sender) request(command string, data url.Values) bool {
	client := &http.Client{}

	url := "https://api.telegram.org/bot" + s.token + "/" + command

	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
