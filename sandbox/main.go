// package main

// import (
// 	"log"
// 	"time"

// 	"tj-system/shared/config"
// 	"tj-system/shared/db"
// 	"tj-system/shared/model"
// )

// func main() {
// 	cfg := config.LoadConfig()
// 	db.Init(cfg.DatabaseURL)

// 	vehicleID := "B1234XYZ"

// 	// 👉 1. Create a simulated payload (as if from MQTT)
// 	payload := &model.VehicleLocationPayload{
// 		VehicleID: vehicleID,
// 		Latitude:  -6.2,
// 		Longitude: 106.816666,
// 		Timestamp: time.Now().Unix(),
// 	}

// 	// 👉 2. Validate payload
// 	if err := payload.Validate(); err != nil {
// 		log.Fatalf("❌ Payload validation failed: %v", err)
// 	}

// 	// 👉 3. Insert into DB
// 	err := db.InsertVehicleLocation(payload)
// 	if err != nil {
// 		log.Fatalf("❌ Failed to insert payload: %v", err)
// 	}

// 	log.Println("✅ Payload successfully inserted into database")

// 	// 👉 4. Query latest location
// 	location, err := db.GetLastVehicleLocation(vehicleID)
// 	if err != nil {
// 		log.Fatalf("❌ Failed to get latest vehicle location: %v", err)
// 	}

// 	log.Printf("📍 Latest location for %s:\nLatitude: %.6f\nLongitude: %.6f\nTimestamp: %d",
// 		location.VehicleID, location.Latitude, location.Longitude, location.Timestamp)
	
// 	// 👉 5. Query location history
// 	history, err := db.GetVehicleLocationHistory(vehicleID, 0, time.Now().Unix())
// 	if err != nil {
// 		log.Fatalf("❌ Failed to get vehicle location history: %v", err)
// 	}

// 	log.Printf("📜 Location history for %s:", vehicleID)
// 	for i, h := range history {
// 		log.Printf("  #%d: Lat: %.6f, Lon: %.6f, Time: %d", i+1, h.Latitude, h.Longitude, h.Timestamp)
// 	}
// }


package main

import (
	"log"
	"github.com/satryo-pramahardi/tj-system/shared/config"
	"github.com/satryo-pramahardi/tj-system/shared/db"

	"github.com/gin-gonic/gin"
	"github.com/satryo-pramahardi/tj-system/backend/api"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg.DatabaseURL)
	log.Println("✅ Database connected")

	r := gin.Default()

	// Route bindings
	r.POST("/vehicle", api.InsertVehicleLocation)
	r.GET("/vehicle/:id", api.GetLatestLocation)
	r.GET("/vehicle/:id/history", api.GetVehicleHistory)

	r.Run(":8080") // 🔥 starts HTTP server on port 8080
}
