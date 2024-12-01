package message

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*DiscordBuilder)(nil)

// DiscordBuilder .
type DiscordBuilder struct{}

// Build .
func (b DiscordBuilder) Build(url, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) string {
	var t1, t2 []string
	for _, t := range team1 {
		t1 = append(t1, fmt.Sprintf("%s %s", t.FirstName, t.LastName))
		if len(t.DiscordUserID) > 0 {
			t1 = append(t1, fmt.Sprintf("<@%s>", t.DiscordUserID))
		}
	}
	for _, t := range team2 {
		t2 = append(t2, fmt.Sprintf("%s %s", t.FirstName, t.LastName))
		if len(t.DiscordUserID) > 0 {
			t2 = append(t2, fmt.Sprintf("<@%s>", t.DiscordUserID))
		}
	}
	return fmt.Sprintf(
		`[Announcement]
%s
%s %s at Table %s
%s 🆚 %s`,
		tName,
		eName,
		ePhase,
		tableNo,
		strings.Join(t1, " "),
		strings.Join(t2, " "),
	)
}
