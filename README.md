# 📖 How to run, endpoints, architecture overview

## Project Overview
This project consists of a backend service, a vehicle simulation script, and a geofence alert worker service. Each component is containerized using Docker.

## How to Run
- Use `docker-compose up` to start the infrastructure stack.
- Navigate to the `backend` directory and build the Docker image using the provided Dockerfile.
- Similarly, build the Docker image for the `geofence_worker`.

## Endpoints
- The backend service provides REST endpoints for retrieving the latest vehicle location and location history.

## Architecture
- The backend service subscribes to MQTT topics and processes geofence logic.
- The geofence worker subscribes to the `geofence_alerts` queue and processes geofence entry messages. 