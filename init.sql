-- initialize database for vehicle_locations and geofence tables
CREATE TABLE IF NOT EXISTS vehicle_locations (
    id SERIAL PRIMARY KEY,
    vehicle_id VARCHAR NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS geofence_events (
  id SERIAL PRIMARY KEY,
  payload JSONB NOT NULL,
  received_at TIMESTAMPTZ DEFAULT NOW()
);