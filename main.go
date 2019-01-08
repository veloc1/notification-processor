package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Configuration struct {
	Token string
}

func main() {
	config := readConfig()

	projects := []Project{
		Project{
			name: "bitbucket-commenter",
			groups: []Group{
				Group{
					name: "developers",
					users: []string{
						"76741468", // veloc1
					},
				},
			},
		},
	}

	handler := &WebhookHandler{
		processors: []Processor{
			&BitbucketProcessor{},
		},
		sender: Sender{
			token: config.Token,
			manager: Manager{
				projects: Projects{
					projects: projects,
				},
			},
		},
	}
	http.Handle("/", handler)

	http.ListenAndServe(":9180", nil)
}

func readConfig() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
