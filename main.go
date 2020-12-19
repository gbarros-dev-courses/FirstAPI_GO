package main

import (
	"net/http"
  "github.com/gin-gonic/gin"

  "github.com/rahmanfadhil/gin-bookstore/models"
  "github.com/rahmanfadhil/gin-bookstore/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}