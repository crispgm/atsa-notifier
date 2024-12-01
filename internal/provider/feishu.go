package provider

import (
	"fmt"
	"net/http"

	"github.com/go-lark/lark"
)

var _ MessageProvider = (*FeishuWebhook)(nil)

// FeishuWebhook .
type FeishuWebhook struct{}

// Send .
func (fw FeishuWebhook) Send(webhookURL string, content []byte) (*http.Response, error) {
	webhook := lark.NewNotificationBot(webhookURL)
	msg := lark.NewMsgBuffer(lark.MsgPost).Text(string(content))
	resp, err := webhook.PostNotificationV2(msg.Build())
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("Lark Error: [%d] %s", resp.Code, resp.Msg)
	}

	return nil, err
}
