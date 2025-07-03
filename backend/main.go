package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/satryo-pramahardi/tj-system/shared/config"
	"github.com/satryo-pramahardi/tj-system/shared/db"
	"github.com/satryo-pramahardi/tj-system/backend/api"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg.DatabaseURL)

	r := gin.Default()

	// url routes
	r.GET("/vehicles/:id/location", api.GetLatestLocation)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
