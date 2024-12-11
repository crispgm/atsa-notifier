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
func (b Discord) CallMatch(template *conf.Template, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	for _, t := range team1 {
		t1 = append(t1, t.Name)
		if len(t.DiscordUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("<@%s>", t.DiscordUserID))
		}
	}
	for _, t := range team2 {
		t2 = append(t2, t.Name)
		if len(t.DiscordUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("<@%s>", t.DiscordUserID))
		}
	}
	data := map[string]interface{}{
		"TournamentName": tName,
		"EventName":      eName,
		"EventPhase":     ePhase,
		"Team1":          strings.Join(t1, template.And),
		"Team2":          strings.Join(t2, template.And),
		"TableNo":        tableNo,
	}
	output, err := EvaluateTemplate("discord_call_match", template.NormalText, data)
	return output, err
}

// RecallPlayer .
func (b Discord) RecallPlayer(template *conf.Template, tName, eName, ePhase, tableNo string, player atsa.Player) (string, error) {
	data := map[string]interface{}{
		"Player":  player.Name,
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("discord_recall_player", template.RecallText, data)
	return output, err
}
