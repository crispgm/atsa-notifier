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
		getLogger(c).Errorln("load players failed")
		ErrorResponse(c, CodeLoadPlayer, "load players failed", nil)
		return ""
	}
	playerDB := atsa.NewPlayerDB(players)

	var team1, team2 []atsa.Player
	for _, player := range params.Team1 {
		team1 = append(team1, *findOrCreatePlayerByID(c, playerDB, &player))
	}
	for _, player := range params.Team2 {
		team2 = append(team2, *findOrCreatePlayerByID(c, playerDB, &player))
	}
	// Create the message content with mention
	template, ok := global.GetGlobalData("templates").(map[string]conf.Template)
	if !ok {
		getLogger(c).Errorln("load templates failed")
		ErrorResponse(c, CodeLoadTemplate, "load templates failed", nil)
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
		getLogger(c).Errorln(templateName, "template not found")
		ErrorResponse(c, CodeLoadTemplate, fmt.Sprintf("[%s] template not found", templateName), nil)
		return ""
	}
	if err != nil {
		getLogger(c).Errorln(err.Error())
		ErrorResponse(c, CodeLoadTemplate, err.Error(), nil)
		return ""
	}
	return strings.TrimSpace(msg)
}

func findOrCreatePlayerByID(c *gin.Context, playerDB *atsa.PlayerDB, player *atsa.Player) *atsa.Player {
	if len(player.ID) > 0 {
		p := playerDB.FindPlayer(player.ID)
		if p != nil {
			return p
		}
		getLogger(c).Warnln("no players found", player.ID)
	}

	newPlayer := atsa.CreatePlayerByFullname(player.Name)
	return &newPlayer
}
