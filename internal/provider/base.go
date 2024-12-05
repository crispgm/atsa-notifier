// Package provider message providers
package provider

import "net/http"

// WebhookMessage represents the structure of the message to send to the Discord webhook
type WebhookMessage struct {
	Content string `json:"content"`
}

// MessageProvider .
type MessageProvider interface {
	Send(webHookURL string, msg *WebhookMessage) (*http.Response, error)
}
