package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"blog-go/internal/database"
	"blog-go/internal/handlers"
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
	router.GET("/posts/:id", handlers.GetPostByID(pool))
	router.POST("/posts", handlers.CreatePost(pool))
	// router.PUT("/posts/:id", handlers.UpdatePost(pool))
	// router.DELETE("/posts/:id", handlers.DeletePost(pool))

	// Run server
	router.Run(":8080")
}
