package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mchernigin/posts-thing/database"
)

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

func getAllAuthors(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authors, err := database.GetAllAuthors(db)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, authors)
	}
}

func getAuthor(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "id is not a number")
			return
		}
		author, err := database.GetAuthorById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, "not found")
			return
		}
		c.JSON(http.StatusOK, author)
	}
}

func createAuthor(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		surname := c.Query("surname")
		website := c.Query("website")
		status := c.Query("status")
		NewAuthor := database.NewAuthor{
			Name:    name,
			Surname: surname,
			Website: website,
			Status:  status,
		}

		authorId, err := database.AddAuthor(db, NewAuthor)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, authorId)
	}
}

func Serve(db *sqlx.DB) {
	router := gin.Default()

	router.GET("/posts", getAllPosts(db))
	router.GET("/posts/:id", getPost(db))
	router.POST("/posts", createPost(db))

	router.GET("/authors", getAllAuthors(db))
	router.GET("/authors/:id", getAuthor(db))
	router.POST("/authors", createAuthor(db))

	router.Run()
}
