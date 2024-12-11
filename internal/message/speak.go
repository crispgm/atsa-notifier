package message

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*Speak)(nil)

// Speak .
type Speak struct{}

// CallMatch .
func (b Speak) CallMatch(template *conf.Template, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) string {
	var t1, t2 []string
	for _, t := range team1 {
		t1 = append(t1, t.Name)
	}
	for _, t := range team2 {
		t2 = append(t2, t.Name)
	}
	return fmt.Sprintf(
		template.NormalSpeak,
		tName,
		eName,
		ePhase,
		strings.Join(t1, " "),
		strings.Join(t2, " "),
		tableNo,
	)
}

// RecallPlayer .
func (b Speak) RecallPlayer(template *conf.Template, tName, eName, ePhase, tableNo string, player atsa.Player) string {
	return fmt.Sprintf(
		template.RecallSpeak,
		player.Name,
		tableNo,
	)
}
