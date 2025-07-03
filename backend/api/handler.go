package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tj-system/shared/db"
)


// GET /vehicle/:id
func GetLatestLocation(c *gin.Context) {
	id := c.Param("id")

	location, err := db.GetLastVehicleLocation(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

// GET /vehicle/:id/history
func GetLocationHistory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle ID"})
		return
	}

	locations, err := db.GetVehicleLocationHistory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch history"})
		return
	}

	c.JSON(http.StatusOK, locations)
}