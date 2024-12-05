package provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-lark/lark"
)

var _ MessageProvider = (*FeishuWebhook)(nil)

// FeishuWebhook .
type FeishuWebhook struct{}

// Send .
func (fw FeishuWebhook) Send(webhookURL string, msg *WebhookMessage) (*http.Response, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}
	webhook := lark.NewNotificationBot(webhookURL)
	mb := lark.NewMsgBuffer(lark.MsgPost).Text(string(content))
	resp, err := webhook.PostNotificationV2(mb.Build())
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("Lark Error: [%d] %s", resp.Code, resp.Msg)
	}

	return nil, err
}
