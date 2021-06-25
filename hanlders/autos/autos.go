package autos

import "github.com/gin-gonic/gin"

func GetAllAutos(c *gin.Context) {
	c.JSON(200, gin.H{
		"autos": "autos",
	})
}
