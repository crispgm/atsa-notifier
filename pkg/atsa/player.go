// Package atsa .
package atsa

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
