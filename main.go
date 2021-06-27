package main

import (
	"github.com/Selopez24/gin-crud/handlers"
	"github.com/Selopez24/gin-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.GET("/bikes", handlers.GetBikes)
	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.CreateBook)
	router.GET("/books/:id", handlers.GetBook)
	router.PATCH("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	router.GET("/status", handlers.GetServerStatus)

	router.Run()
}
