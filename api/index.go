// package api
package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/crispgm/atsa-notifier/internal/conf"
	"github.com/crispgm/atsa-notifier/internal/global"
	"github.com/crispgm/atsa-notifier/internal/handler"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	path := "./conf/conf.yml"
	cfg, err := conf.LoadConf(path)
	if err != nil {
		panic(err)
	}
	// Load players
	global.LoadGlobalData(cfg)

	// Init web handlers
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	r := gin.New()

	// Init log
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	r.Use(func(c *gin.Context) {
		c.Set("logger", log)
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.WithFields(logrus.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"client_ip":   c.ClientIP(),
			"remote_ip":   c.RemoteIP(),
			"status":      c.Writer.Status(),
			"response_ms": duration.Milliseconds(),
			"response_us": duration.Microseconds(),
		}).Info("handled request")
	})

	// Serve static files
	r.Static("/css", "./web/css")
	r.Static("/js", "./web/js")
	r.Static("/img", "./web/img")

	r.StaticFile("/", "web/index.html")

	// Define routes
	r.GET("/sync", handler.SyncHandler)
	r.POST("/notify", handler.NotifyHandler)

	// Run the server
	if len(cfg.Mode) > 0 {
		gin.SetMode(cfg.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.ServeHTTP(w, req)
}
