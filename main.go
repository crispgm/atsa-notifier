// Package main ATSA Notifier
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// WebhookMessage represents the structure of the message to send to the Discord webhook
type WebhookMessage struct {
	Content string `json:"content"`
}

func main() {
	// Replace with your Discord webhook URL
	webhookURL := "https://discord.com/api/webhooks/1311930779034714156/fKxhXxEVSGi7J-kp5rf8FrmUiF8XoylmSS-BreVujFp9dOAM0xrZrdgPsODR6UeCNnvj"

	// Replace with the user's ID you want to mention
	userID := "674619029415264334"

	// Create the message content with mention
	message := WebhookMessage{
		Content: fmt.Sprintf(
			`[Announcement]
%s <@%s> 🆚 Harrod HO
%s %s at Table %d`,
			"David Zhang",
			userID,
			"Open Single", // Event
			"Qualification",
			3, // Table
		),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Send POST request to the Discord webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Failed to send message, status code:", resp.StatusCode)
	// 	return
	// }

	fmt.Println("Message sent successfully!")
}
