package handler

import (
	"fmt"

	"github.com/crispgm/atsa-notifier/internal/global"
	"github.com/crispgm/atsa-notifier/internal/scraper"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
	"github.com/gin-gonic/gin"
)

// SyncParams .
type SyncParams struct {
	URL string `form:"url" binding:"required"`
}

// SyncMatch .
type SyncMatch struct {
	Team1   []atsa.Player `json:"team1"`
	Team2   []atsa.Player `json:"team2"`
	TableNo string        `json:"tableNo"`
}

// SyncOutput .
type SyncOutput struct {
	URL     string      `json:"url"`
	Page    string      `json:"page"`
	Matches []SyncMatch `json:"matches"`
}

// SyncHandler serves the main page.
func SyncHandler(c *gin.Context) {
	var params SyncParams
	err := c.BindQuery(&params)
	if err != nil {
		ErrorResponse(c, CodeParamsErr, err.Error(), nil)
		return
	}
	s := scraper.NewScraper()
	matches, err := s.Scrape(params.URL)
	if err != nil {
		ErrorResponse(c, CodeParamsErr, err.Error(), nil)
		return
	}
	var output SyncOutput
	players, ok := global.GetGlobalData("players").([]atsa.Player)
	if !ok {
		ErrorResponse(c, CodeLoadPlayer, "load player failed", nil)
		return
	}
	playerDB := atsa.NewPlayerDB(players)
	for _, match := range *matches {
		matchWithPlayerInfo := SyncMatch{
			TableNo: match.TableNo,
		}
		for _, name := range match.Team1 {
			matchWithPlayerInfo.Team1 = append(matchWithPlayerInfo.Team1, *findOrCreatePlayerByName(playerDB, name))
		}
		for _, name := range match.Team2 {
			matchWithPlayerInfo.Team2 = append(matchWithPlayerInfo.Team2, *findOrCreatePlayerByName(playerDB, name))
		}
		output.Matches = append(output.Matches, matchWithPlayerInfo)
	}
	output.URL = params.URL
	output.Page = "kickertool_live"
	SuccessResponse(c, output)
}

func findOrCreatePlayerByName(playerDB *atsa.PlayerDB, name string) *atsa.Player {
	p := playerDB.FindPlayersByFullName(name)
	np := len(p)
	if np == 1 {
		return &(p[0])
	}

	if np > 1 {
		fmt.Println("multiple players found", name)
	} else {

		fmt.Println("no players found", name)
	}

	newPlayer := atsa.CreatePlayerByFullname(name)
	return &newPlayer
}
