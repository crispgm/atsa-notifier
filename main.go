// Package main ATSA Notifier
package main

import (
	"fmt"

	"github.com/crispgm/atsa-notifier/internal/announcer"
	"github.com/crispgm/atsa-notifier/internal/message"
	"github.com/crispgm/atsa-notifier/internal/provider"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
	"github.com/hegedustibor/htgo-tts/voices"
)

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
	msg := &provider.WebhookMessage{
		Content: discordBuilder.Build(webhookURL, tournamentName, eventName, eventPhase, tableNo, team1, team2),
	}
	// Send POST request to the Discord webhook
	discord := provider.DiscordWebhook{}
	_, err := discord.Send(webhookURL, msg)
	if err != nil {
		panic(err)
	}

	// Create text message for speech
	ab := message.AnnouncementBuilder{}
	announcement := ab.Build("", tournamentName, eventName, eventPhase, tableNo, team1, team2)
	announcer.TextToSpeech(announcement, voices.English)

	fmt.Println("Message sent successfully!")
}
