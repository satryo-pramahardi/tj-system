package api

import {
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"tj-system/shared/db"
	"tj-system/shared/model"
}


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