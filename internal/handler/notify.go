package handler

import (
	"github.com/crispgm/atsa-notifier/internal/provider"
	"github.com/gin-gonic/gin"
)

// NotifyParams .
type NotifyParams struct {
	TournamentName string   `json:"tournamentName"`
	EventName      string   `json:"eventName"`
	EventPhase     string   `json:"eventPhase"`
	Team1          []string `json:"team1"`
	Team2          []string `json:"team2"`
	TableNo        string   `json:"tableNo"`
	Locale         string   `json:"locale"`

	Text string `json:"text"`

	DiscordWebhookURL string `json:"discordWebhookURL"`
	FeishuWebhookURL  string `json:"feishuWebhookURL"`
}

// NotifyOutput .
type NotifyOutput struct {
	MsgType string `json:"msgType"`
	Text    string `json:"text"`
}

// NotifyHandler serves the main page.
func NotifyHandler(c *gin.Context) {
	var params NotifyParams
	err := c.BindJSON(&params)
	if err != nil {
		ErrorResponse(c, CodeParamsErr, err.Error(), nil)
		return
	}

	var msgType string
	if len(params.DiscordWebhookURL) > 0 {
		msgType = "discord"
	} else if len(params.FeishuWebhookURL) > 0 {
		msgType = "feishu"
	} else if len(params.Text) > 0 {
		msgType = "text"
	} else {
		msgType = "announcement"
	}

	var msg provider.WebhookMessage
	if len(params.Text) > 0 {
		msg.Content = params.Text
	} else {
		msg.Content = buildMessage(c, &params, msgType)
	}

	var output NotifyOutput
	output.MsgType = msgType
	output.Text = msg.Content
	// Send POST request to the Discord webhook
	if msgType == "discord" {
		discordSender := provider.DiscordWebhook{}
		_, err = discordSender.Send(params.DiscordWebhookURL, &msg)
		if err != nil {
			ErrorResponse(c, CodeParamsErr, err.Error(), nil)
			return
		}
	} else if msgType == "feishu" {
		// do nothing rn
	} else if msgType == "announcement" {
		// do nothing rn
	}

	SuccessResponse(c, output)
}
