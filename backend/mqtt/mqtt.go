package mqtt

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"tj-system/shared/db"
	"tj-system/shared/model"
)

// StartMQTT initializes the MQTT client, connects to the broker,
// and subscribes to the given topic.
func StartMQTT(brokerURL, topic string, handler func(v *model.VehicleLocationPayload) error) {
	
	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID("transjakarta-mqtt-subscriber").
		SetAutoReconnect(true)

	opts.OnConnect = func(c mqtt.Client) {
		log.Printf("Connected to MQTT broker at %s", brokerURL)

		if token := c.Subscribe(topic, 1, handleMessage); token.Wait() && token.Error() != nil {
			log.Printf("Failed to subscribe to topic %s: %v", topic, token.Error())
		} else {
			log.Printf("Subscribed to topic: %s", topic)
		}
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection failed: %v", token.Error())
	}
}

// handleMessage is the callback for processing incoming messages.
var handleMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message on topic: %s", msg.Topic())

	var payload model.VehicleLocationPayload
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Printf("Invalid JSON payload: %v", err)
		return
	}

	if err := payload.Validate(); err != nil {
		log.Printf("Payload validation failed: %v", err)
		return
	}

	if err := db.InsertVehicleLocation(&payload); err != nil {
		log.Printf("Failed to insert payload into database: %v", err)
		return
	}

	log.Printf("Stored location for vehicle %s", payload.VehicleID)
}
