package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBikes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"bikes": "bikes",
	})
}
