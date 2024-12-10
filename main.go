// Package main ATSA Notifier
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/internal/message"
	"github.com/crispgm/atsa-notifier/internal/provider"
	"github.com/crispgm/atsa-notifier/internal/scraper"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

func main() {
	args := os.Args
	path := ""
	if len(args) < 2 {
		path = "./conf/conf.yml"
	}
	// Init
	cfg, err := conf.LoadConf(path)
	if err != nil {
		panic(err)
	}
	// Load players
	dbpath := cfg.ATSADB.DefaultPath
	players, err := conf.LoadPlayerLocalDB(dbpath)
	if err != nil {
		panic(err)
	}
	playerDB := atsa.NewPlayerDB(players)
	// Load tournament info
	liveURL := "https://live.kickertool.de/crispfoosball/tournaments/IiGr5-u1QMG3sXbzLrXEt/live"
	webhookURL := "https://discord.com/api/webhooks/1311930779034714156/fKxhXxEVSGi7J-kp5rf8FrmUiF8XoylmSS-BreVujFp9dOAM0xrZrdgPsODR6UeCNnvj"
	tournamentName := "ATSA50 YangShengCup"
	eventName := "Open Single"
	eventPhase := "Qualification"

	// Scraping
	scr := scraper.NewScraper()
	matches, err := scr.Scrape(liveURL)
	if err != nil {
		panic(err)
	}

	discordBuilder := message.DiscordBuilder{}
	discordSender := provider.DiscordWebhook{}
	if matches != nil && len(*matches) > 0 {
		for _, match := range *matches {
			var team1, team2 []atsa.Player
			for _, player := range match.Team1 {
				p := playerDB.FindPlayers(player)
				var pName = player
				if len(p) == 1 {
					pName = p[0].FullName
				} else {
					fmt.Println("cannot find", player)
				}
				team1 = append(team1, convertPlayer(pName))
			}
			for _, player := range match.Team2 {
				p := playerDB.FindPlayers(player)
				var pName = player
				if len(p) == 1 {
					pName = p[0].FullName
				} else {
					fmt.Println("cannot find", player)
				}
				team2 = append(team2, convertPlayer(pName))
			}
			// Create the message content with mention
			msg := &provider.WebhookMessage{
				Content: discordBuilder.Build(webhookURL, tournamentName, eventName, eventPhase, match.TableNo, team1, team2),
			}
			// Send POST request to the Discord webhook
			_, err := discordSender.Send(webhookURL, msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Message sent successfully!")
}

func convertPlayer(fullName string) atsa.Player {
	var firstName, lastName string
	names := strings.Split(fullName, " ")
	if len(names) == 2 {
		firstName = names[0]
		lastName = names[1]
	} else if len(names) > 2 {
		firstName = strings.Join(names[0:len(names)-1], " ")
		lastName = names[len(names)-1]
	} else {
		lastName = names[0]
	}

	return atsa.Player{FirstName: firstName, LastName: lastName}
}
