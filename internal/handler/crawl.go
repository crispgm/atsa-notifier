package handler

import (
	"github.com/crispgm/atsa-notifier/internal/scraper"
	"github.com/gin-gonic/gin"
)

// CrawlParams .
type CrawlParams struct {
	URL string `form:"url" binding:"required"`
}

// CrawlOutput .
type CrawlOutput struct {
	URL     string          `json:"url"`
	Page    string          `json:"page"`
	Matches []scraper.Match `json:"matches"`
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
	output.URL = params.URL
	output.Page = "kickertool_live"
	output.Matches = *matches
	SuccessResponse(c, output)
}
