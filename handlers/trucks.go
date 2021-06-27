package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTrucks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"trucks": "trucks",
	})
}
