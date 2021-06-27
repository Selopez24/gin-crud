package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAutos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"autos": "autos",
	})
}
