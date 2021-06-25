package trucks

import "github.com/gin-gonic/gin"

func GetAllTrucks(c *gin.Context) {
	c.JSON(200, gin.H{
		"trucks": "trucks",
	})
}
