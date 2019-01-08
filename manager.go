package main

type Manager struct {
	projects Projects
}

func (m Manager) GetReceivers(data NotifyData) []string {
	for _, p := range m.projects.All() {
		if p.name == data.project {
			return getReceiversFromGroups(data.groups, p.groups)
		}
	}
	return []string{}
}

func getReceiversFromGroups(groupNames []string, groups []Group) []string {
	result := []string{}
	for _, groupName := range groupNames {
		for _, group := range groups {
			if groupName == group.name {
				result = append(result, group.users...)
			}
		}
	}

	return result
}
