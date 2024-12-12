// Package main web entrance
package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/internal/global"
	"github.com/crispgm/atsa-notifier/internal/handler"
)

func main() {
	// Init
	args := os.Args
	path := ""
	if len(args) < 2 {
		path = "./conf/conf.yml"
	}
	cfg, err := conf.LoadConf(path)
	if err != nil {
		panic(err)
	}
	// Load players
	global.LoadGlobalData(cfg)

	// web handlers
	r := gin.Default()

	// Serve static files
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.Static("/img", "./web/img")

	r.StaticFile("/", "web/index.html")

	// Define routes
	r.GET("/crawl", handler.CrawlHandler)
	r.POST("/notify", handler.NotifyHandler)

	// Run the server
	if len(cfg.Mode) > 0 {
		gin.SetMode(cfg.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Run(cfg.Port)
}
