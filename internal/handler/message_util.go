package handler

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/internal/global"
	"github.com/crispgm/atsa-notifier/internal/message"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
	"github.com/gin-gonic/gin"
)

func buildMessage(
	c *gin.Context,
	params *NotifyParams,
	msgType string,
) string {
	players, ok := global.GetGlobalData("players").([]atsa.Player)
	if !ok {
		ErrorResponse(c, CodeLoadPlayer, "load player failed", nil)
		return ""
	}
	playerDB := atsa.NewPlayerDB(players)

	var team1, team2 []atsa.Player
	for _, player := range params.Team1 {
		p := playerDB.FindPlayers(player)
		var pName = player
		if len(p) == 1 {
			pName = p[0].Name
		} else {
			fmt.Println("cannot find", player)
		}
		team1 = append(team1, convertPlayer(pName))
	}
	for _, player := range params.Team2 {
		p := playerDB.FindPlayers(player)
		var pName = player
		if len(p) == 1 {
			pName = p[0].Name
		} else {
			fmt.Println("cannot find", player)
		}
		team2 = append(team2, convertPlayer(pName))
	}
	// Create the message content with mention
	template, ok := global.GetGlobalData("templates").(map[string]conf.Template)
	if !ok {
		ErrorResponse(c, CodeLoadTemplate, "load template failed", nil)
		return ""
	}
	if params.Locale == "" {
		params.Locale = "enUS"
	}
	var msg string
	if template, ok := template[params.Locale]; ok {
		if msgType == "discord" {
			discordBuilder := message.DiscordBuilder{}
			msg = discordBuilder.Build(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1, team2)
		} else if msgType == "announcement" {
			announcementBuilder := message.AnnouncementBuilder{}
			msg = announcementBuilder.Build(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1, team2)
		}
	} else {
		ErrorResponse(c, CodeLoadTemplate, fmt.Sprintf("[%s] template not found", params.Locale), nil)
		return ""
	}
	return strings.TrimSpace(msg)
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
