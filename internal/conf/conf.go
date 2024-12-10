// Package conf .
package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf .
type Conf struct {
	Locales []string     `yaml:"locales"`
	ATSADB  ATSADatabase `yaml:"atsa_database"`
}

// ATSADatabase .
type ATSADatabase struct {
	DefaultPath string `yaml:"default_path"`
}

// LoadConf .
func LoadConf(path string) (*Conf, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Conf
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
