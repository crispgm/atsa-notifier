// Package main ATSA Notifier
package main

import (
	"encoding/json"
	"fmt"

	"github.com/crispgm/atsa-notifier/internal/message"
	"github.com/crispgm/atsa-notifier/internal/provider"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

// WebhookMessage represents the structure of the message to send to the Discord webhook
type WebhookMessage struct {
	Content string `json:"content"`
}

func main() {
	// Load tournament info
	webhookURL := "https://discord.com/api/webhooks/1311930779034714156/fKxhXxEVSGi7J-kp5rf8FrmUiF8XoylmSS-BreVujFp9dOAM0xrZrdgPsODR6UeCNnvj"
	tournamentName := "ATSA50 YangShengCup"
	eventName := "Open Single"
	eventPhase := "Qualification"
	tableNo := "3"
	team1 := []atsa.Player{
		{
			FirstName:     "David",
			LastName:      "Zhang",
			DiscordUserID: "674619029415264334",
		},
	}
	team2 := []atsa.Player{
		{
			FirstName: "Harrod",
			LastName:  "HO",
		},
	}
	// Create the message content with mention
	discordBuilder := message.DiscordBuilder{}
	msg := WebhookMessage{
		Content: discordBuilder.Build(webhookURL, tournamentName, eventName, eventPhase, tableNo, team1, team2),
	}
	content, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Send POST request to the Discord webhook
	discord := provider.DiscordWebhook{}
	_, err = discord.Send(webhookURL, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Message sent successfully!")
}
