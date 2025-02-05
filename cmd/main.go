package main

import (
	"blog-go/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	//connect to db

	db, err := database.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Blog!"})
	})

	router.Run(":8080")
}
