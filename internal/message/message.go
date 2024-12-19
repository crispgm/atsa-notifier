// Package message .
package message

import (
	"bytes"
	"text/template"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

// Builder interface for Message Builder
type Builder interface {
	CallMatch(
		template *conf.Template,
		prefix string,
		postfix string,
		tableNo string,
		team1 []atsa.Player,
		team2 []atsa.Player,
	) (string, error)
	RecallPlayer(
		template *conf.Template,
		prefix string,
		postfix string,
		tableNo string,
		player atsa.Player,
	) (string, error)
}

// EvaluateTemplate .
func EvaluateTemplate(name, tpl string, data any) (string, error) {
	instance, err := template.New(name).Parse(tpl)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	instance.Execute(buf, data)

	return buf.String(), nil
}
