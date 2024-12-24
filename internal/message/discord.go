package message

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*Discord)(nil)

// Discord .
type Discord struct{}

// CallMatch .
func (b Discord) CallMatch(template *conf.Template, msgParams *MsgParams, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	nameOpt := &atsa.NameOpt{
		Native: msgParams.Native,
	}
	for _, t := range team1 {
		name := t.OptName(nameOpt)
		if len(t.DiscordUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("%s <@%s>", name, t.DiscordUserID))
		} else {
			t1 = append(t1, name)
		}
	}
	for _, t := range team2 {
		name := t.OptName(nameOpt)
		if len(t.DiscordUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("%s <@%s>", name, t.DiscordUserID))
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
	output, err := EvaluateTemplate("discord_call_match", template.NormalText, data)
	return output, err
}

// RecallPlayer .
func (b Discord) RecallPlayer(template *conf.Template, msgParams *MsgParams, player atsa.Player) (string, error) {
	playerName := player.OptName(&atsa.NameOpt{Native: msgParams.Native})
	if len(player.DiscordUserID) > 0 {
		playerName = fmt.Sprintf("%s <@%s>", playerName, player.DiscordUserID)
	}
	data := map[string]interface{}{
		"Player":  playerName,
		"TableNo": msgParams.TableNo,
	}
	output, err := EvaluateTemplate("discord_recall_player", template.RecallText, data)
	return output, err
}
