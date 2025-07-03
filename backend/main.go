package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"tj-system/shared/config"
	"tj-system/shared/db"
	"tj-system/backend/api"
)

func main() {
	// Load config and init DB
	cfg := config.LoadConfig()
	db.Init(cfg.DatabaseURL)
	log.Println("✅ Database connected")

	// Start Gin router
	r := gin.Default()

	// ✅ Set your URL routes here
	r.GET("/vehicles/:id/location", api.GetLatestLocation)
	r.GET("/vehicles/:id/history", api.GetVehicleHistory)

	// Start server
	r.Run(":8080") // Now server listens on localhost:8080
}
