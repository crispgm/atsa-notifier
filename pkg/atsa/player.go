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

// NameOpt .
type NameOpt struct {
	Native bool
}

// OptName select the correct name with options
func (p Player) OptName(opt *NameOpt) string {
	if opt.Native && len(p.NativeName) > 0 {
		return p.NativeName
	}

	return p.Name
}

// CreatePlayerByFullname creates a player with full name
func CreatePlayerByFullname(fullName string) Player {
	var firstName, lastName, name string
	names := strings.Split(fullName, " ")
	if len(names) == 2 {
		firstName = names[0]
		lastName = names[1]
		name = fmt.Sprintf("%s %s", firstName, lastName)
	} else if len(names) > 2 {
		firstName = strings.Join(names[0:len(names)-1], " ")
		lastName = names[len(names)-1]
		name = fmt.Sprintf("%s %s", firstName, lastName)
	} else {
		lastName = names[0]
		name = lastName
	}

	return Player{
		FullName:  fullName,
		FirstName: firstName,
		LastName:  lastName,
		Name:      name,
	}
}
