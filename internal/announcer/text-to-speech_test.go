package announcer

import (
	"testing"

	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/stretchr/testify/assert"
)

func TestTTS(t *testing.T) {
	assert.NoError(t, TextToSpeech("hello, world", voices.English))
}
