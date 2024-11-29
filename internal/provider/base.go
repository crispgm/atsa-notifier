// Package provider message providers
package provider

import "net/http"

// MessageProvider .
type MessageProvider interface {
	Send(webHookURL string, content []byte) (*http.Response, error)
}
