package handlers

import (
	"blog-go/internal/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPosts(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := pool.Query(context.Background(), "SELECT id, title, content, author_id FROM posts")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch posts: " + err.Error(),
			})
			return
		}
		defer rows.Close()

		var posts []models.Post
		for rows.Next() {
			var p models.Post
			err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to scan post: " + err.Error(),
				})
				return
			}

			posts = append(posts, p)
		}

		c.JSON(http.StatusOK, posts)
	}
}

func CreatePost(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request: " + err.Error(),
			})
			return
		}
		_, err := pool.Exec(
			context.Background(),
			"INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)",
			post.Title, post.Content, post.AuthorID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create post: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, post)
	}
}
