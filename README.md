# Transjakarta Vehicle Tracking System

This project is a backend system for tracking Transjakarta vehicles in real-time. It collects location data from vehicles via MQTT, stores it in a PostgreSQL database, and provides a RESTful API to query the latest position or travel history.

It also includes geofencing support—when a vehicle enters a predefined area, an event is triggered and published through RabbitMQ. Everything runs inside Docker for easy setup and isolation.

This system is part of a take-home assignment, but it's built with real-world patterns and scalability in mind.

## How to Run This Project

This system is containerized using Docker Compose. To start all services:

```bash
docker compose up --build -d
```

## Project Status

This project is a work-in-progress. Here's what has been completed so far, and what still needs to be implemented:

### ✅ Completed

- 🧠 **API Server (Gin)**
  - Up and running on port `8080`.
  - Handles requests and returns proper JSON responses.

- 📦 **CRU for Vehicle Location**
  - Create and Read operations via MQTT and REST API.
  - Location history query with UNIX timestamp filters.

- 🔊 **MQTT Server Subscription**
  - Listens to topics like `/fleet/vehicle/+/location`.
  - Parses and **inserts location data into PostgreSQL** — *logic written, but not fully tested*.

- 🧪 **Publisher Simulator (WIP)**
  - CLI-based Go script to simulate vehicle location data.
  - Works locally, but **publishing to broker is still buggy** (being fixed).

---

### 🔧 In Progress / Not Yet Implemented

- 📍 **Geofencing Detection**
  - Logic for detecting entry into defined geofenced zones is **not yet implemented**.
  - No RabbitMQ events are triggered yet.

- 📤 **RabbitMQ Integration**
  - Event queue publishing will begin after geofence detection is implemented.

## Simulating Vehicle Movement with MQTT Publisher

To test the system, you can simulate vehicle location updates by running the MQTT publisher included in this project.

This script will:
- Generate fake GPS coordinates (can be randomized or scripted).
- Publish to the MQTT topic `/fleet/vehicle/<vehicle_id>/location`.
- Send JSON payloads like: `{"lat": -6.2, "lng": 106.8}` at a set interval.

---

### Simulation (Publisher) Guide

1. **Open a terminal**

   Navigate to the root of the project (or anywhere you have the publisher code).

2. **Run the simulation with:**

   ```bash
   go run ./publisher --vehicle-id TJ001 --count 50 --trip-length 200
   ```

    CLI Flags Explained:
    --vehicle-id TJ001
    The unique ID for the simulated vehicle. This will be used in the MQTT topic /fleet/vehicle/TJ001/location.

    --count 50
    The number of ticks or location updates to send. Each tick is one MQTT message.

    --trip-length 200
    The total length of the simulated route (in arbitrary distance units).
    Once the vehicle reaches the end of this trip, it will turn around and head back to simulate a round trip.

## API Reference

### GET `/location/:id`

Get the **latest known location** of a vehicle.

**Example:**
GET /vehicle/TJ001location

**Response:**
```json
{
  "vehicle_id": "TJ001",
  "lat": -6.2,
  "lng": 106.8,
  "timestamp": 1720073100
}
```

### GET `/location/:id/history?start=<unix>&end=<unix>`

Get **location history** for a specific vehicle between two UNIX timestamps.

**Query Parameters:**
- `start` – Start time in **UNIX timestamp** format.
- `end` – End time in **UNIX timestamp** format.

**Example:**
GET /location/TJ001/history?start=1720070000&end=1720079999

**Response:**
```json
[
  {
    "lat": -6.2,
    "lng": 106.8,
    "timestamp": 1720073100
  },
  {
    "lat": -6.3,
    "lng": 106.81,
    "timestamp": 1720073160
  }
]
```

## Envvars
Create new `.env.docker` file in root of the project with following variables:

```
MQTT_BROKER=tcp://localhost:1883
MQTT_TOPIC=/fleet/vehicle/+/location
DB_URL=postgres://user:password@db:5432/tracking?sslmode=disable
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
```

## Design Consideration
There are 3 main modules:
- Backend: work as core service, hosted API server & MQTT
- Geofence Worker: worker to host queues, will recieve alerts in geofence events
- Publisher: non-dockerized function to simulate vehicle

Tables:
- vehicle_location: to store vehicle locations
- geofence_events: simple DB with JSONDB column, for flexibilty and future dev