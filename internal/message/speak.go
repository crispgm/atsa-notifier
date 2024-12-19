package message

import (
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*Speak)(nil)

// Speak .
type Speak struct{}

// CallMatch .
func (b Speak) CallMatch(template *conf.Template, prefix, postfix, tableNo string, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	for _, t := range team1 {
		t1 = append(t1, t.Name)
	}
	for _, t := range team2 {
		t2 = append(t2, t.Name)
	}
	data := map[string]interface{}{
		"Prefix":  prefix,
		"Postfix": postfix,
		"Team1":   strings.Join(t1, template.And),
		"Team2":   strings.Join(t2, template.And),
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("speak_call_match", template.NormalSpeak, data)
	return output, err
}

// RecallPlayer .
func (b Speak) RecallPlayer(template *conf.Template, prefix, postfix, tableNo string, player atsa.Player) (string, error) {
	data := map[string]interface{}{
		"Player":  player.Name,
		"TableNo": tableNo,
	}
	output, err := EvaluateTemplate("speak_recall_player", template.RecallSpeak, data)
	return output, err
}
