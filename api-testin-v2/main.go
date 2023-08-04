package main

import (
	"github.com/gin-gonic/gin"

	"example/api-testin-v2/controllers"
	"example/api-testin-v2/models"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}

func main() {
	r := SetupRouter()
	models.ConnectDatabase()
	r.Run()
}
