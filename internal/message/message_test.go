package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageTemplate(t *testing.T) {
	output, err := EvaluateTemplate("testPass",
		"hello {{.Name}}, please go {{.Action}}",
		map[string]interface{}{
			"Name":   "Amanda",
			"Action": "hiking",
		})
	if assert.NoError(t, err) {
		assert.Equal(t, "hello Amanda, please go hiking", output)
	}
	output, err = EvaluateTemplate("testFail", "Hello{{a}}", nil)
	assert.Error(t, err)
	assert.Empty(t, output)
}
