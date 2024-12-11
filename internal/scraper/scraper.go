// Package scraper scrapes kickertool
package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

// Scraper .
type Scraper struct {
	batch  int
	result [][]*Match
}

// NewScraper .
func NewScraper() *Scraper {
	return &Scraper{
		batch:  0,
		result: make([][]*Match, 0),
	}
}

// Scrape .
func (s *Scraper) Scrape(url string) (*[]Match, error) {
	var matches []Match
	c := colly.NewCollector()
	c.OnHTML("div.live-match-row", func(e *colly.HTMLElement) {
		m := Match{}
		e.ForEach("div.table-name", func(idx int, e *colly.HTMLElement) {
			m.TableNo = e.Text
		})
		e.ForEach("div.time", func(idx int, e *colly.HTMLElement) {
			m.Duration = e.Text
		})
		e.ForEach("div.participant", func(idx int, e *colly.HTMLElement) {
			e.ForEach("div.participant-names.left", func(idx int, e *colly.HTMLElement) {
				m.Team1 = splitNames(e.Text)
			})
			e.ForEach("div.participant-names.right", func(idx int, e *colly.HTMLElement) {
				m.Team2 = splitNames(e.Text)
			})
		})
		matches = append(matches, m)
	})
	c.Visit(url)

	return &matches, nil
}

func splitNames(text string) []string {
	var names []string
	if strings.Contains(text, "/") {
		names = strings.Split(text, "/")
		names[0] = strings.TrimSpace(names[0])
		names[1] = strings.TrimSpace(names[1])
	} else {
		names = append(names, text)
	}
	return names
}

// LastResult .
func (s *Scraper) LastResult() []*Match {
	if s.batch <= 0 || s.batch > len(s.result) {
		return nil
	}
	return s.result[s.batch-1]
}

// Results .
func (s Scraper) Results() [][]*Match {
	return s.result
}
