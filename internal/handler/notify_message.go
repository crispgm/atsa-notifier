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
	templateName string,
) string {
	players, ok := global.GetGlobalData("players").([]atsa.Player)
	if !ok {
		ErrorResponse(c, CodeLoadPlayer, "load player failed", nil)
		return ""
	}
	playerDB := atsa.NewPlayerDB(players)

	var team1, team2 []atsa.Player
	for _, player := range params.Team1 {
		p := playerDB.FindPlayer(player)
		if p != nil {
			team1 = append(team1, *p)
		} else {
			team1 = append(team1, convertPlayer(player))
			fmt.Println("no players found", player)
		}
	}
	for _, player := range params.Team2 {
		p := playerDB.FindPlayer(player)
		if p != nil {
			team2 = append(team2, *p)
		} else {
			team2 = append(team2, convertPlayer(player))
			fmt.Println("no players found", player)
		}
	}
	// Create the message content with mention
	template, ok := global.GetGlobalData("templates").(map[string]conf.Template)
	if !ok {
		ErrorResponse(c, CodeLoadTemplate, "load template failed", nil)
		return ""
	}
	if params.Locale == "" {
		params.Locale = "en-US"
	}
	var msg string
	var err error
	if template, ok := template[params.Locale]; ok {
		speakBuilder := message.Speak{}
		discordBuilder := message.Discord{}
		feishuBuilder := message.Feishu{}
		if msgType == "speak" {
			if templateName == "call_match" {
				msg, err = speakBuilder.CallMatch(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1, team2)
			} else if templateName == "recall_player" {
				msg, err = speakBuilder.RecallPlayer(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1[0])
			}
		} else if msgType == "discord" {
			if templateName == "call_match" {
				msg, err = discordBuilder.CallMatch(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1, team2)
			} else if templateName == "recall_player" {
				msg, err = discordBuilder.RecallPlayer(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1[0])
			}
		} else if msgType == "feishu" {
			if templateName == "call_match" {
				msg, err = feishuBuilder.CallMatch(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1, team2)
			} else if templateName == "recall_player" {
				msg, err = feishuBuilder.RecallPlayer(&template, params.TournamentName, params.EventName, params.EventPhase, params.TableNo, team1[0])
			}
		}
	} else {
		ErrorResponse(c, CodeLoadTemplate, fmt.Sprintf("[%s] template not found", params.Locale), nil)
		return ""
	}
	if err != nil {
		ErrorResponse(c, CodeLoadTemplate, err.Error(), nil)
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

	return atsa.Player{
		FullName:  fullName,
		FirstName: firstName,
		LastName:  lastName,
		Name:      fmt.Sprintf("%s %s", firstName, lastName),
	}
}
