package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satryo-pramahardi/tj-system/shared/db"
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