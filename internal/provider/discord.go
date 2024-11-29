package provider

import (
	"bytes"
	"fmt"
	"net/http"
)

var _ DiscordWebhook

// DiscordWebhook .
type DiscordWebhook struct{}

// Send .
func (dw DiscordWebhook) Send(webhookURL string, content []byte) (*http.Response, error) {
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	return resp, err
}
