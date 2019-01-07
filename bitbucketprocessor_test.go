package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCanHandle(t *testing.T) {
	processor := BitbucketProcessor{}

	result := processor.canHandle("bitbucket")
	if !result {
		t.Errorf("BitbucketProcessor can't handle appropriate service string")
	}

	result = processor.canHandle("bitbuck")
	if result {
		t.Errorf("BitbucketProcessor can handle inappropriate service string")
	}
}

func TestEmptyData(t *testing.T) {
	processor := BitbucketProcessor{}

	req := httptest.NewRequest("GET", "/", nil)

	_, err := processor.process(req)
	if err == nil {
		t.Errorf("BitbucketProcessor can process empty event")
	}
}

func TestValidData(t *testing.T) {
	processor := BitbucketProcessor{}

	body := `{
		"pullrequest": {
		  "rendered": {
			"description": {
			  "raw": "",
			  "markup": "markdown",
			  "html": "",
			  "type": "rendered"
			},
			"title": {
			  "raw": "main.go edited online with Bitbucket",
			  "markup": "markdown",
			  "html": "<p>main.go edited online with Bitbucket</p>",
			  "type": "rendered"
			}
		  },
		  "type": "pullrequest",
		  "description": "",
		  "links": {
			"decline": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/decline"
			},
			"commits": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/commits"
			},
			"self": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3"
			},
			"comments": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/comments"
			},
			"merge": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/merge"
			},
			"html": {
			  "href": "https://bitbucket.org/veloc1/bitbucket-commenter/pull-requests/3"
			},
			"activity": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/activity"
			},
			"diff": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/diff"
			},
			"approve": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/approve"
			},
			"statuses": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/pullrequests/3/statuses"
			}
		  },
		  "title": "main.go edited online with Bitbucket",
		  "close_source_branch": false,
		  "reviewers": [],
		  "id": 3,
		  "destination": {
			"commit": {
			  "hash": "45f9818e43ff",
			  "type": "commit",
			  "links": {
				"self": {
				  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/commit/45f9818e43ff"
				},
				"html": {
				  "href": "https://bitbucket.org/veloc1/bitbucket-commenter/commits/45f9818e43ff"
				}
			  }
			},
			"repository": {
			  "links": {
				"self": {
				  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter"
				},
				"html": {
				  "href": "https://bitbucket.org/veloc1/bitbucket-commenter"
				},
				"avatar": {
				  "href": "https://bytebucket.org/ravatar/%7B87a6d9c2-2403-44ef-8ad9-fd936ac45227%7D?ts=go"
				}
			  },
			  "type": "repository",
			  "name": "bitbucket-commenter",
			  "full_name": "veloc1/bitbucket-commenter",
			  "uuid": "{87a6d9c2-2403-44ef-8ad9-fd936ac45227}"
			},
			"branch": {
			  "name": "master"
			}
		  },
		  "created_on": "2019-01-07T21:13:08.182500+00:00",
		  "summary": {
			"raw": "",
			"markup": "markdown",
			"html": "",
			"type": "rendered"
		  },
		  "source": {
			"commit": {
			  "hash": "8015474dee7d",
			  "type": "commit",
			  "links": {
				"self": {
				  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter/commit/8015474dee7d"
				},
				"html": {
				  "href": "https://bitbucket.org/veloc1/bitbucket-commenter/commits/8015474dee7d"
				}
			  }
			},
			"repository": {
			  "links": {
				"self": {
				  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter"
				},
				"html": {
				  "href": "https://bitbucket.org/veloc1/bitbucket-commenter"
				},
				"avatar": {
				  "href": "https://bytebucket.org/ravatar/%7B87a6d9c2-2403-44ef-8ad9-fd936ac45227%7D?ts=go"
				}
			  },
			  "type": "repository",
			  "name": "bitbucket-commenter",
			  "full_name": "veloc1/bitbucket-commenter",
			  "uuid": "{87a6d9c2-2403-44ef-8ad9-fd936ac45227}"
			},
			"branch": {
			  "name": "testing-webhooks"
			}
		  },
		  "comment_count": 0,
		  "state": "OPEN",
		  "task_count": 0,
		  "participants": [],
		  "reason": "",
		  "updated_on": "2019-01-07T21:13:08.216897+00:00",
		  "author": {
			"username": "veloc1",
			"display_name": "Pavel Bartashuk",
			"account_id": "557058:658ff884-0b9d-4f5e-99f4-06ed3b9ec737",
			"links": {
			  "self": {
				"href": "https://api.bitbucket.org/2.0/users/veloc1"
			  },
			  "html": {
				"href": "https://bitbucket.org/veloc1/"
			  },
			  "avatar": {
				"href": "https://bitbucket.org/account/veloc1/avatar/"
			  }
			},
			"nickname": "veloc1",
			"type": "user",
			"uuid": "{3b2c0a02-79b9-4f7c-8363-50f626a96ea5}"
		  },
		  "merge_commit": null,
		  "closed_by": null
		},
		"repository": {
		  "scm": "git",
		  "website": "",
		  "name": "bitbucket-commenter",
		  "links": {
			"self": {
			  "href": "https://api.bitbucket.org/2.0/repositories/veloc1/bitbucket-commenter"
			},
			"html": {
			  "href": "https://bitbucket.org/veloc1/bitbucket-commenter"
			},
			"avatar": {
			  "href": "https://bytebucket.org/ravatar/%7B87a6d9c2-2403-44ef-8ad9-fd936ac45227%7D?ts=go"
			}
		  },
		  "full_name": "veloc1/bitbucket-commenter",
		  "owner": {
			"username": "veloc1",
			"display_name": "Pavel Bartashuk",
			"account_id": "557058:658ff884-0b9d-4f5e-99f4-06ed3b9ec737",
			"links": {
			  "self": {
				"href": "https://api.bitbucket.org/2.0/users/veloc1"
			  },
			  "html": {
				"href": "https://bitbucket.org/veloc1/"
			  },
			  "avatar": {
				"href": "https://bitbucket.org/account/veloc1/avatar/"
			  }
			},
			"nickname": "veloc1",
			"type": "user",
			"uuid": "{3b2c0a02-79b9-4f7c-8363-50f626a96ea5}"
		  },
		  "type": "repository",
		  "is_private": false,
		  "uuid": "{87a6d9c2-2403-44ef-8ad9-fd936ac45227}"
		},
		"actor": {
		  "username": "veloc1",
		  "display_name": "Pavel Bartashuk",
		  "account_id": "557058:658ff884-0b9d-4f5e-99f4-06ed3b9ec737",
		  "links": {
			"self": {
			  "href": "https://api.bitbucket.org/2.0/users/veloc1"
			},
			"html": {
			  "href": "https://bitbucket.org/veloc1/"
			},
			"avatar": {
			  "href": "https://bitbucket.org/account/veloc1/avatar/"
			}
		  },
		  "nickname": "veloc1",
		  "type": "user",
		  "uuid": "{3b2c0a02-79b9-4f7c-8363-50f626a96ea5}"
		}
	  }`

	req := httptest.NewRequest("POST", "/", strings.NewReader(body))
	req.Header.Add("X-Event-Key", "pullrequest:created")
	req.Header.Add("User-Agent", "Bitbucket-Webhooks/2.0")
	req.Header.Add("Content-Type", "application/json")

	data, err := processor.process(req)
	if err != nil {
		t.Errorf("BitbucketProcessor cannot process valid event")
	}

	expected := "Pullrequest created at bitbucket-commenter"
	if data.message != expected {
		t.Errorf("Message incorrect, got: %s, want: %s.", data.message, expected)
	}

	expected = "bitbucket-commenter"
	if data.project != expected {
		t.Errorf("Project incorrect, got: %s, want: %s.", data.project, expected)
	}
}
