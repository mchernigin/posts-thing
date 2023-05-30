package routes

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mchernigin/posts-thing/database"
)

func SetupPostsRoutes(router *gin.Engine, db *sqlx.DB) {
	postsRoutes := router.Group("/posts")

	postsRoutes.GET("/", getAllPosts(db))
	postsRoutes.POST("/", createPost(db))
	postsRoutes.GET("/:id", getPost(db))
	postsRoutes.DELETE("/:id", deletePost(db))
}

func getAllPosts(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := database.GetAllPosts(db)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, posts)
	}
}

func getPost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "id is not a number")
			return
		}
		post, err := database.GetPostById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, "not found")
			return
		}
		c.JSON(http.StatusOK, post)
	}
}

func deletePost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "id is not a number")
			return
		}
		err = database.DeletePostById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, "not found")
			return
		}
		c.Status(http.StatusOK)
	}
}

func createPost(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		content := c.Query("content")
		newPost := database.NewPost{
			Title:            title,
			Content:          content,
			PublishTimestamp: time.Now(),
		}

		postId, err := database.AddPost(db, newPost)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, postId)
	}
}
