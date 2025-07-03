package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type VehicleLocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

// metersToLatOffset converts meters to approx latitude degrees
func metersToLatOffset(m float64) float64 {
	return m / 111320 // 1 degree = 111320 meters
}

func main() {
	var (
		broker     = flag.String("broker", "tcp://localhost:1883", "MQTT broker URL")
		vehicleID  = flag.String("vehicle-id", "TJ001", "Vehicle ID")
		interval   = flag.Int("interval", 2, "Seconds between messages")
		count      = flag.Int("count", 0, "Number of messages to send (0=infinite)")
		tripLength = flag.Float64("trip-length", 200, "Trip length in meters before turning")
	)
	flag.Parse()

	client := mqtt.NewClient(mqtt.NewClientOptions().
		AddBroker(*broker).
		SetClientID("publisher-" + *vehicleID).
		SetAutoReconnect(true),
	)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to broker: %v", token.Error())
	}

	topic := fmt.Sprintf("/fleet/vehicle/%s/location", *vehicleID)

	baseLat := -6.2000
	baseLon := 106.816666
	step := metersToLatOffset(5) // move ~5 meters per tick
	offset := 0.0
	direction := 1.0
	limit := metersToLatOffset(*tripLength)

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()

	sent := 0
	for {
		lat := baseLat + offset
		payload := VehicleLocationPayload{
			VehicleID: *vehicleID,
			Latitude:  lat,
			Longitude: baseLon,
			Timestamp: time.Now().Unix(),
		}

		data, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Marshal error: %v", err)
			continue
		}

		token := client.Publish(topic, 1, false, data)
		token.Wait()
		if token.Error() != nil {
			log.Printf("Publish error: %v", token.Error())
		} else {
			log.Printf("Published: lat=%.6f dir=%.0f", lat, direction)
		}

		offset += direction * step
		if math.Abs(offset) >= limit {
			direction *= -1
			log.Printf("Turned around at lat=%.6f", lat)
		}

		sent++
		if *count > 0 && sent >= *count {
			break
		}

		<-ticker.C
	}

	log.Printf("Finished sending %d messages", sent)
}
