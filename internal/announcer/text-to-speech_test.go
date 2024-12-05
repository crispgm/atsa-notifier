package announcer

import (
	"testing"

	"github.com/hegedustibor/htgo-tts/voices"
)

func TestTTSEnglish(t *testing.T) {
	TextToSpeech("Open Single Qualification Albert Lee versus David Beckham at Table 3", voices.English)
}

func TestTTSChinese(t *testing.T) {
	TextToSpeech("Open Single 预选赛 李选手 对阵 David Beckham 球桌3", voices.Chinese)
}
