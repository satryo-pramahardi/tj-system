package model

import (
	"errors"
	"fmt"
)

// VehicleLocationPayload represents the structure of the vehicle location
// received via MQTT and stored in the database.
type VehicleLocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"` // UNIX epoch seconds
}

// Validate checks that all required fields in the payload are present and valid.
func (v *VehicleLocationPayload) Validate() error {
	if v.VehicleID == "" {
		return errors.New("vehicle_id is required")
	}
	if v.Latitude < -90 || v.Latitude > 90 {
		return fmt.Errorf("latitude %.6f out of bounds", v.Latitude)
	}
	if v.Longitude < -180 || v.Longitude > 180 {
		return fmt.Errorf("longitude %.6f out of bounds", v.Longitude)
	}
	if v.Timestamp <= 0 {
		return fmt.Errorf("invalid timestamp: %d", v.Timestamp)
	}
	return nil
}

// ToDBValues returns the values ready to be used in DB insert.
func (v *VehicleLocationPayload) ToDBValues() []interface{} {
	return []interface{}{
		v.VehicleID,
		v.Latitude,
		v.Longitude,
		v.Timestamp,
	}
}
