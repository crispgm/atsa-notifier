package message

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*Feishu)(nil)

// Feishu .
type Feishu struct{}

// CallMatch .
func (b Feishu) CallMatch(template *conf.Template, msgParams *MsgParams, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	nameOpt := &atsa.NameOpt{
		Native: msgParams.Native,
	}
	for _, t := range team1 {
		name := t.OptName(nameOpt)
		if len(t.FeishuUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", name, t.FeishuUserID))
		} else {
			t1 = append(t1, name)
		}
	}
	for _, t := range team2 {
		name := t.OptName(nameOpt)
		if len(t.FeishuUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", name, t.FeishuUserID))
		} else {
			t2 = append(t2, name)
		}
	}
	data := map[string]interface{}{
		"Prefix":  msgParams.Prefix,
		"Postfix": msgParams.Postfix,
		"TableNo": msgParams.TableNo,
		"Team1":   strings.Join(t1, template.And),
		"Team2":   strings.Join(t2, template.And),
	}
	output, err := EvaluateTemplate("feishu_call_match", template.NormalText, data)
	return output, err
}

// RecallPlayer .
func (b Feishu) RecallPlayer(template *conf.Template, msgParams *MsgParams, player atsa.Player) (string, error) {
	playerName := player.OptName(&atsa.NameOpt{Native: msgParams.Native})
	if len(player.FeishuUserID) > 0 {
		playerName = fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", playerName, player.FeishuUserID)
	}
	data := map[string]interface{}{
		"Player":  playerName,
		"TableNo": msgParams.TableNo,
	}
	output, err := EvaluateTemplate("feishu_recall_player", template.RecallText, data)
	return output, err
}
