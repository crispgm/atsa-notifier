package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var _ MessageProvider = (*DiscordWebhook)(nil)

// DiscordWebhook .
type DiscordWebhook struct{}

// Send .
func (dw DiscordWebhook) Send(webhookURL string, msg *WebhookMessage) (*http.Response, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	return resp, err
}
