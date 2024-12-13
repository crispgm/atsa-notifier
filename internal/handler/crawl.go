package handler

import (
	"fmt"

	"github.com/crispgm/atsa-notifier/internal/global"
	"github.com/crispgm/atsa-notifier/internal/scraper"
	"github.com/crispgm/atsa-notifier/pkg/atsa"
	"github.com/gin-gonic/gin"
)

// CrawlParams .
type CrawlParams struct {
	URL string `form:"url" binding:"required"`
}

// CrawlMatch .
type CrawlMatch struct {
	Team1   []atsa.Player `json:"team1"`
	Team2   []atsa.Player `json:"team2"`
	TableNo string        `json:"tableNo"`
}

// CrawlOutput .
type CrawlOutput struct {
	URL     string       `json:"url"`
	Page    string       `json:"page"`
	Matches []CrawlMatch `json:"matches"`
}

// CrawlHandler serves the main page.
func CrawlHandler(c *gin.Context) {
	var params CrawlParams
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
	var output CrawlOutput
	players, ok := global.GetGlobalData("players").([]atsa.Player)
	if !ok {
		ErrorResponse(c, CodeLoadPlayer, "load player failed", nil)
		return
	}
	playerDB := atsa.NewPlayerDB(players)
	for _, match := range *matches {
		matchWithPlayerInfo := CrawlMatch{
			TableNo: match.TableNo,
		}
		for _, name := range match.Team1 {
			p := playerDB.FindPlayersByFullName(name)
			if len(p) == 1 {
				matchWithPlayerInfo.Team1 = append(matchWithPlayerInfo.Team1, p[0])
			} else {
				matchWithPlayerInfo.Team1 = append(matchWithPlayerInfo.Team1, atsa.CreatePlayerByFullname(name))
				fmt.Println("no or multiple players found", name)
			}
		}
		for _, name := range match.Team2 {
			p := playerDB.FindPlayersByFullName(name)
			if len(p) == 1 {
				matchWithPlayerInfo.Team2 = append(matchWithPlayerInfo.Team2, p[0])
			} else {
				matchWithPlayerInfo.Team2 = append(matchWithPlayerInfo.Team2, atsa.CreatePlayerByFullname(name))
				fmt.Println("no or multiple players found", name)
			}
		}
		output.Matches = append(output.Matches, matchWithPlayerInfo)
	}
	output.URL = params.URL
	output.Page = "kickertool_live"
	SuccessResponse(c, output)
}
