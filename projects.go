package main

type Group struct {
	name  string
	users []string
}

type Project struct {
	name   string
	groups []Group
}

type Projects struct {
	projects []Project
}

func (p Projects) All() []Project {
	return p.projects
}
