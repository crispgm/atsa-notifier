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

	MsgType  string `json:"msgType"`  // speak, discord, feishu
	Template string `json:"template"` // e.g. call_match, recall_player or text; text is given by user manually
	Text     string `json:"text"`

	DiscordWebhookURL string `json:"discordWebhookURL"`
	FeishuWebhookURL  string `json:"feishuWebhookURL"`
}

// NotifyOutput .
type NotifyOutput struct {
	MsgType  string `json:"msgType"`
	Template string `json:"template"`
	Text     string `json:"text"`
}

// NotifyHandler serves the main page.
func NotifyHandler(c *gin.Context) {
	var params NotifyParams
	err := c.BindJSON(&params)
	if err != nil {
		ErrorResponse(c, CodeParamsErr, err.Error(), nil)
		return
	}

	if params.MsgType == "discord" {
		if len(params.DiscordWebhookURL) == 0 {
			ErrorResponse(c, CodeParamsErr, "discordWebhookURL is not set", nil)
			return
		}
	} else if params.MsgType == "feishu" {
		if len(params.FeishuWebhookURL) == 0 {
			ErrorResponse(c, CodeParamsErr, "feishuWebhookURL is not set", nil)
			return
		}
	} else {
		params.MsgType = "speak"
	}

	var msg provider.WebhookMessage
	if len(params.Text) > 0 {
		msg.Content = params.Text
	} else {
		if !(params.Template == "call_match" || params.Template == "recall_player") {
			ErrorResponse(c, CodeParamsErr, "Template is not existed", nil)
			return
		}
		msg.Content = buildMessage(c, &params, params.MsgType, params.Template)
	}

	var output NotifyOutput
	output.MsgType = params.MsgType
	output.Template = params.Template
	output.Text = msg.Content
	// Send POST request to the Discord webhook
	if output.MsgType == "discord" {
		discordSender := provider.DiscordWebhook{}
		_, err = discordSender.Send(params.DiscordWebhookURL, &msg)
		if err != nil {
			ErrorResponse(c, CodeParamsErr, err.Error(), nil)
			return
		}
	} else if output.MsgType == "feishu" {
		feishuSender := provider.FeishuWebhook{}
		_, err = feishuSender.Send(params.FeishuWebhookURL, &msg)
		if err != nil {
			ErrorResponse(c, CodeParamsErr, err.Error(), nil)
			return
		}
	} else if output.MsgType == "speak" {
		// nothing needed to be done
	}

	SuccessResponse(c, output)
}
