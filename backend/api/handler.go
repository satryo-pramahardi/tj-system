package api

import (
	"net/http"
	"strconv"

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
	id := c.Param("id")
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start timestamp"})
		return
	}
	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end timestamp"})
		return
	}

	locations, err := db.GetVehicleLocationHistory(id, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch history"})
		return
	}

	c.JSON(http.StatusOK, locations)
}
