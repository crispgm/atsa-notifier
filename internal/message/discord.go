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
func (b Discord) CallMatch(template *conf.Template, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) string {
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
	return fmt.Sprintf(
		template.NormalText,
		tName,
		eName,
		ePhase,
		strings.Join(t1, " & "),
		strings.Join(t2, " & "),
		tableNo,
	)
}

// RecallPlayer .
func (b Discord) RecallPlayer(template *conf.Template, tName, eName, ePhase, tableNo string, player atsa.Player) string {
	return fmt.Sprintf(
		template.RecallText,
		player.Name,
		tableNo,
	)
}
