package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tj-system/shared/config"
	"tj-system/shared/db"
	"tj-system/backend/api"
	"tj-system/backend/mqtt"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg.DatabaseURL)

	r := gin.Default()

	// mqtt routes
	go mqtt.StartMQTT(cfg.MQTTBroker, cfg.MQTTTopic, db.InsertVehicleLocation)

	// url routes
	r.GET("/vehicles/:id/location", api.GetLatestLocation)
	r.GET("/vehicles/:id/history", api.GetLocationHistory)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run("0.0.0.0:8080")
}
