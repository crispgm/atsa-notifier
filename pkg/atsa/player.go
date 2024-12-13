// Package atsa .
package atsa

import (
	"fmt"
	"strings"
)

// Player an ATSA player
type Player struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	Name        string `json:"name"`
	NativeName  string `json:"nativeName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	ITSFNumber  string `json:"itsfNumber"`
	CountryCode string `json:"countryCode"`
	CountryFlag string `json:"countryFlag"`
	ProfilePic  string `json:"profilePic"`
	TourLog     string `json:"tourLog"`
	QRCode      string `json:"qrCode"`
	ProPic      string `json:"proPic"`

	DiscordUserID string `json:"discordUserID"`
	FeishuUserID  string `json:"feishuUserID"`
}

// CreatePlayerByFullname creates a player with full name
func CreatePlayerByFullname(fullName string) Player {
	var firstName, lastName string
	names := strings.Split(fullName, " ")
	if len(names) == 2 {
		firstName = names[0]
		lastName = names[1]
	} else if len(names) > 2 {
		firstName = strings.Join(names[0:len(names)-1], " ")
		lastName = names[len(names)-1]
	} else {
		lastName = names[0]
	}

	return Player{
		FullName:  fullName,
		FirstName: firstName,
		LastName:  lastName,
		Name:      fmt.Sprintf("%s %s", firstName, lastName),
	}
}
