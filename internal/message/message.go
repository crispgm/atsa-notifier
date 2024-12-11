// Package message .
package message

import (
	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

// Builder interface for Message Builder
type Builder interface {
	CallMatch(
		template *conf.Template,
		tName string,
		eName string,
		ePhase string,
		tableNo string,
		team1 []atsa.Player,
		team2 []atsa.Player,
	) string
	RecallPlayer(
		template *conf.Template,
		tName string,
		eName string,
		ePhase string,
		tableNo string,
		player atsa.Player,
	) string
}
