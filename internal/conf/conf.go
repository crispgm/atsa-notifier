// Package conf .
package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf .
type Conf struct {
	Port      string              `yaml:"port"`
	Templates map[string]Template `yaml:"templates"`
	ATSADB    ATSADatabase        `yaml:"atsa_database"`
}

// Template .
type Template struct {
	NormalSpeak string `yaml:"normal_speak"`
	NormalText  string `yaml:"normal_text"`
	RecallSpeak string `yaml:"recall_speak"`
	RecallText  string `yaml:"recall_text"`
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
