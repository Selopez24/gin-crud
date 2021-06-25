package bikes

import "github.com/gin-gonic/gin"

func GetBikes(c *gin.Context) {
	c.JSON(200, gin.H{
		"bikes": "bikes",
	})
}
