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
func (b Discord) CallMatch(template *conf.Template, prefix, postfix, tableNo string, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	for _, t := range team1 {
		if len(t.DiscordUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("%s <@%s>", t.Name, t.DiscordUserID))
		} else {
			t1 = append(t1, t.Name)
		}
	}
	for _, t := range team2 {
		if len(t.DiscordUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("%s <@%s>", t.Name, t.DiscordUserID))
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
	output, err := EvaluateTemplate("discord_call_match", template.NormalText, data)
	return output, err
}

// RecallPlayer .
func (b Discord) RecallPlayer(template *conf.Template, prefix, postfix, tableNo string, player atsa.Player) (string, error) {
	playerName := player.Name
	if len(player.DiscordUserID) > 0 {
		playerName = fmt.Sprintf("%s <@%s>", playerName, player.DiscordUserID)
	}
	data := map[string]interface{}{
		"Player":  playerName,
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("discord_recall_player", template.RecallText, data)
	return output, err
}
