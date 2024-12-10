package conf

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/crispgm/atsa-notifier/pkg/atsa"
)

// LoadPlayerLocalDB .
func LoadPlayerLocalDB(path string) ([]atsa.Player, error) {
	var players []atsa.Player
	var err error

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return players, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return players, err
	}

	for _, record := range records {
		var player atsa.Player
		player.FullName = record[0]
		player.Name = record[1]
		player.ID = record[2]
		player.NativeName = record[3]
		player.FirstName = record[4]
		player.LastName = record[5]
		player.Gender = record[6]
		player.ITSFNumber = record[7]
		player.CountryCode = record[8]
		player.CountryFlag = record[9]
		player.ProfilePic = record[10]
		player.TourLog = record[11]
		player.QRCode = record[12]
		player.ProPic = record[13]
		players = append(players, player)
	}

	return players, nil
}
