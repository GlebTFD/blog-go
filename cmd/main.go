package main

import (
	"blog-go/internal/database"
	"blog-go/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to db
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// Initializing the router
	router := gin.Default()

	// Registering routers
	router.GET("/posts", handlers.GetPosts(pool))
	router.POST("/posts", handlers.CreatePost(pool))

	//Server run
	router.Run(":8080")
}
