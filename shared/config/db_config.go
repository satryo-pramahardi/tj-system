package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	DatabaseURL string // ← Computed here

	MQTTBroker string
	MQTTTopic  string
}

func LoadConfig() AppConfig {
	env := os.Getenv("GO_ENV")

	if env == "" || env != "docker" {
		if err := godotenv.Load(".env.dev"); err != nil {
			fmt.Println("Error loading .env.dev file")
		} else {
			fmt.Println("Loaded .env.dev file")
			fmt.Println("DATABASE_HOST:", os.Getenv("DATABASE_HOST"))
			// Print other variables as needed
		}
	}

	host := getEnv("DATABASE_HOST", "localhost")
	port := getEnv("DATABASE_PORT", "5432")
	user := getEnv("DATABASE_USER", "admin")
	pass := getEnv("DATABASE_PASSWORD", "password")
	name := getEnv("DATABASE_NAME", "transjakarta")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, name,
	)

	return AppConfig{
		DBHost:      host,
		DBPort:      port,
		DBUser:      user,
		DBPassword:  pass,
		DBName:      name,
		DatabaseURL: dbURL,

		MQTTBroker: getEnv("MQTT_BROKER", "tcp://localhost:1883"),
		MQTTTopic:  getEnv("MQTT_TOPIC", "/fleet/vehicle/+/location"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
