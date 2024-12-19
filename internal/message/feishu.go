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
func (b Feishu) CallMatch(template *conf.Template, prefix, postfix, tableNo string, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	for _, t := range team1 {
		if len(t.FeishuUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", t.Name, t.FeishuUserID))
		} else {
			t1 = append(t1, t.Name)
		}
	}
	for _, t := range team2 {
		if len(t.FeishuUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", t.Name, t.FeishuUserID))
		} else {
			t2 = append(t2, t.Name)
		}
	}
	data := map[string]interface{}{
		"Prefix":  prefix,
		"Postfix": postfix,
		"Team1":   strings.Join(t1, template.And),
		"Team2":   strings.Join(t2, template.And),
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("feishu_call_match", template.NormalText, data)
	return output, err
}

// RecallPlayer .
func (b Feishu) RecallPlayer(template *conf.Template, prefix, postfix, tableNo string, player atsa.Player) (string, error) {
	playerName := player.Name
	if len(player.FeishuUserID) > 0 {
		playerName = fmt.Sprintf("%s <at user_id=\"%s\">@user</at>", playerName, player.FeishuUserID)
	}
	data := map[string]interface{}{
		"Player":  playerName,
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("feishu_recall_player", template.RecallText, data)
	return output, err
}
