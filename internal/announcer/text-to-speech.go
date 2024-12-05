// Package announcer implements text to speech
package announcer

import (
	"github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
)

// TextToSpeech .
func TextToSpeech(text string, language string) error {
	speech := htgotts.Speech{
		Folder:   "/tmp/audio",
		Language: language,
		Handler:  &handlers.Native{},
	}
	return speech.Speak(text)
}
