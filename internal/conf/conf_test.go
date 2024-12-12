package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	path := "../../conf/conf.yml"
	if os.Getenv("CI") == "true" {
	}
	conf, err := LoadConf(path)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, conf.Mode)
		assert.NotEmpty(t, conf.Port)
		assert.NotEmpty(t, conf.ATSADB.LocalPath)
		assert.Empty(t, conf.ATSADB.WebURL)
		assert.NotEmpty(t, conf.Templates)
	}
}
