package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	s := NewScraper()
	assert.Zero(t, s.batch)
	assert.Empty(t, s.result)

	s.Scrape("https://live.kickertool.de/crispfoosball/tournaments/IiGr5-u1QMG3sXbzLrXEt/live")
}
