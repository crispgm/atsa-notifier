package message

import (
	"fmt"
	"strings"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

var _ Builder = (*AnnouncementBuilder)(nil)

// AnnouncementBuilder .
type AnnouncementBuilder struct{}

// Build .
func (b AnnouncementBuilder) Build(template *conf.Template, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) string {
	var t1, t2 []string
	for _, t := range team1 {
		t1 = append(t1, fmt.Sprintf("%s %s", t.FirstName, t.LastName))
	}
	for _, t := range team2 {
		t2 = append(t2, fmt.Sprintf("%s %s", t.FirstName, t.LastName))
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
