package scraper

// Match .
type Match struct {
	Team1    []string `json:"team1,omitempty"`
	Team2    []string `json:"team2,omitempty"`
	TableNo  string   `json:"tableNo,omitempty"`
	Duration string   `json:"duration,omitempty"`
}
