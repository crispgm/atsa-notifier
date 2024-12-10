package atsa

import (
	"strings"
)

// PlayerDB ATSA player DB utility
type PlayerDB struct {
	db    []Player
	cache map[string][]Player
}

// NewPlayerDB .
func NewPlayerDB(players []Player) *PlayerDB {
	for i, p := range players {
		players[i].FullName = strings.TrimSpace(p.FullName)
		players[i].Name = strings.TrimSpace(p.Name)
		players[i].NativeName = strings.TrimSpace(p.NativeName)
	}

	return &PlayerDB{
		db:    players,
		cache: make(map[string][]Player),
	}
}

// FindPlayersByFullName find player by full name
// No other names
// No fuzzy
func (db *PlayerDB) FindPlayersByFullName(input string) []Player {
	var players []Player
	// Normalize
	input = strings.TrimSpace(input)
	if input == "" {
		return players
	}

	if res, ok := db.cache[input]; ok {
		return res
	}
	for _, p := range db.db {
		if p.FullName == input {
			players = append(players, p)
		}
	}
	db.cache[input] = players

	return players
}

// FindPlayers find player by full name, native name and name
func (db *PlayerDB) FindPlayers(input string) []Player {
	var players []Player
	// Normalize
	input = strings.TrimSpace(input)
	if input == "" {
		return players
	}

	if res, ok := db.cache[input]; ok {
		return res
	}
	for _, p := range db.db {
		if p.FullName == input || p.NativeName == input || p.Name == input {
			players = append(players, p)
		}
	}
	db.cache[input] = players

	return players
}
