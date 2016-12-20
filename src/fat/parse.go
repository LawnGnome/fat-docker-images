package fat

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Architecture struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
}

type Repository struct {
	Architectures []Architecture `yaml:"architectures"`
	Target        string         `yaml:"target"`
}

func ParseRepositoriesFile(file string) ([]Repository, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	if err := yaml.Unmarshal(data, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
