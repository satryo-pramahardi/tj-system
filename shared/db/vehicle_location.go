package db

import (
	"log"
	"tj-system/shared/model"

	_ "github.com/lib/pq"
)

// InsertVehicleLocation inserts a new vehicle location into the database.
func InsertVehicleLocation(v *model.VehicleLocationPayload) error {
	query := `
		INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
		VALUES ($1, $2, $3, to_timestamp($4))
	`
	_, err := DB.Exec(query, v.ToDBValues()...)
	if err != nil {
		log.Printf("InsertVehicleLocation error: %v", err)
	}
	return err
}

// GetLastVehicleLocation returns the most recent location for a given vehicle.
func GetLastVehicleLocation(vehicleID string) (*model.VehicleLocationPayload, error) {
	query := `
		SELECT vehicle_id, latitude, longitude, FLOOR(EXTRACT(EPOCH FROM timestamp))::BIGINT AS timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`
	row := DB.QueryRow(query, vehicleID)

	var v model.VehicleLocationPayload
	err := row.Scan(&v.VehicleID, &v.Latitude, &v.Longitude, &v.Timestamp)
	if err != nil {
		log.Printf("GetLastVehicleLocation error: %v", err)
		return nil, err
	}
	return &v, nil
}

// GetVehicleLocationHistory returns all vehicle locations between start and end UNIX timestamps.
func GetVehicleLocationHistory(vehicleID string, start, end int64) ([]model.VehicleLocationPayload, error) {
	query := `
		SELECT vehicle_id, latitude, longitude, FLOOR(EXTRACT(EPOCH FROM timestamp))::BIGINT AS timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		AND timestamp BETWEEN to_timestamp($2) AND to_timestamp($3)
		ORDER BY timestamp ASC

	`
	rows, err := DB.Query(query, vehicleID, start, end)
	if err != nil {
		log.Printf("GetVehicleLocationHistory query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var history []model.VehicleLocationPayload
	for rows.Next() {
		var v model.VehicleLocationPayload
		if err := rows.Scan(&v.VehicleID, &v.Latitude, &v.Longitude, &v.Timestamp); err != nil {
			log.Printf("Row scan error: %v", err)
			continue
		}
		history = append(history, v)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	return history, nil
}
