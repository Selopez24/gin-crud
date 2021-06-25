package main

import (
	bikes_hanlder "gin-crud/hanlders/bikes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/bikes", bikes_hanlder.GetBikes)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
