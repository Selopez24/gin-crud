package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerStatus(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
