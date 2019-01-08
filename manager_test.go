package main

import (
	"testing"
)

func TestGetReceivers(t *testing.T) {
	projects := []Project{
		Project{
			name: "bitbucket-commenter",
			groups: []Group{
				Group{
					name: "developers",
					users: []string{
						"76741468",
					},
				},
			},
		},
	}
	manager := &Manager{
		projects: Projects{
			projects: projects,
		},
	}

	result := manager.GetReceivers(NotifyData{
		project: "wrong-project",
		groups: []string{
			"developers",
		},
	})

	if len(result) != 0 {
		t.Errorf("Manager get receivers, but it shouldn't")
	}

	result = manager.GetReceivers(NotifyData{
		project: "bitbucket-commenter",
		groups: []string{
			"wrong-group",
		},
	})

	if len(result) != 0 {
		t.Errorf("Manager get receivers, but it shouldn't")
	}

	result = manager.GetReceivers(NotifyData{
		project: "bitbucket-commenter",
		groups: []string{
			"developers",
		},
	})

	if len(result) == 0 {
		t.Errorf("Manager don't get receivers")
	}

	if result[0] != "76741468" {
		t.Errorf("Manager get wrong receiver, got: %s, want: %s.", result[0], "76741468")
	}
}
