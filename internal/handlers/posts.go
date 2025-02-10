package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"blog-go/internal/models"
)

func GetPosts(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := pool.Query(
			context.Background(),
			"SELECT id, title, content, author_id FROM posts",
		)
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

func GetPostByID(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post models.Post

		err := pool.QueryRow(context.Background(), "SELECT id, title, content, author_id FROM posts WHERE id = $1", id).
			Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to fetch post: " + err.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

func UpdatePost(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post models.Post

		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request: " + err.Error(),
			})
			return
		}

		_, err := pool.Exec(
			context.Background(),
			"UPDATE posts SET title = $1, content = $2 WHERE id = $3",
			post.Title,
			post.Content,
			id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update post: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, post)
	}
}

func DeletePost(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := pool.Exec(
			context.Background(),
			"DELETE FROM posts WHERE id = $1",
			id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete post: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
	}
}
