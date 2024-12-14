package scraper

// Match .
type Match struct {
	Team1    []string `json:"team1"`
	Team2    []string `json:"team2"`
	TableNo  string   `json:"tableNo"`
	Duration string   `json:"duration"`
	Valid    bool     `json:"valid"`
}
