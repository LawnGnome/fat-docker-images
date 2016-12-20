package fat

import (
	"github.com/heroku/docker-registry-client/registry"
)

var hub *registry.Registry

func EnumerateTags(source string) ([]string, error) {
	if hub == nil {
		var err error

		hub, err = registry.New("https://registry-1.docker.io/", "", "")
		if err != nil {
			return nil, err
		}
	}

	return hub.Tags(source)
}

func (r Repository) TagsInAllArchitectures() ([]string, error) {
	var tags []string
	tagCounts := make(map[string]int)

	for _, arch := range r.Architectures {
		archTags, err := EnumerateTags(arch.Source)
		if err != nil {
			return nil, err
		}

		for _, tag := range archTags {
			tagCounts[tag] = tagCounts[tag] + 1
		}
	}

	for tag, count := range tagCounts {
		if count == len(r.Architectures) {
			tags = append(tags, tag)
		}
	}

	return tags, nil
}
