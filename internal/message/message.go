// Package message .
package message

import "github.com/crispgm/atsa-notifier/pkg/atsa"

// Builder interface for Message Builder
type Builder interface {
	Build(url, tName, eName, ePhase, tableNo string, team1 []atsa.Player, team2 []atsa.Player) string
}
