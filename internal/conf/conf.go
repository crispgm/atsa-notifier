// Package conf .
package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf .
type Conf struct {
	Port      string              `yaml:"port"`
	Mode      string              `yaml:"mode"`
	ATSADB    ATSADatabase        `yaml:"atsa_database"`
	Templates map[string]Template `yaml:"templates"`
}

// Template .
type Template struct {
	And         string `yaml:"and"`
	NormalSpeak string `yaml:"normal_speak"`
	NormalText  string `yaml:"normal_text"`
	RecallSpeak string `yaml:"recall_speak"`
	RecallText  string `yaml:"recall_text"`
}

// ATSADatabase .
type ATSADatabase struct {
	LocalPath string `yaml:"local_path"`
	WebURL    string `yaml:"web_url"`
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
