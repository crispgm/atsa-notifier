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
func (b Speak) CallMatch(template *conf.Template, msgParams *MsgParams, team1 []atsa.Player, team2 []atsa.Player) (string, error) {
	var t1, t2 []string
	nameOpt := &atsa.NameOpt{
		Native: msgParams.Native,
	}
	for _, t := range team1 {
		t1 = append(t1, t.OptName(nameOpt))
	}
	for _, t := range team2 {
		t2 = append(t2, t.OptName(nameOpt))
	}
	data := map[string]interface{}{
		"Prefix":  msgParams.Prefix,
		"Postfix": msgParams.Postfix,
		"TableNo": msgParams.TableNo,
		"Team1":   strings.Join(t1, template.And),
		"Team2":   strings.Join(t2, template.And),
	}
	output, err := EvaluateTemplate("speak_call_match", template.NormalSpeak, data)
	return output, err
}

// RecallPlayer .
func (b Speak) RecallPlayer(template *conf.Template, msgParams *MsgParams, player atsa.Player) (string, error) {
	nameOpt := &atsa.NameOpt{
		Native: msgParams.Native,
	}
	data := map[string]interface{}{
		"Player":  player.OptName(nameOpt),
		"TableNo": msgParams.TableNo,
	}
	output, err := EvaluateTemplate("speak_recall_player", template.RecallSpeak, data)
	return output, err
}
