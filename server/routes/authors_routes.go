package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mchernigin/posts-thing/database"
)

func SetupAuthorsRoutes(router *gin.Engine, db *sqlx.DB) {
	authorsRoutes := router.Group("/authors")

	authorsRoutes.GET("/", getAllAuthors(db))
	authorsRoutes.POST("/", createAuthor(db))
	authorsRoutes.GET("/:id", getAuthor(db))
	authorsRoutes.DELETE("/:id", deleteAuthor(db))
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

func deleteAuthor(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "id is not a number")
			return
		}
		err = database.DeleteAuthorById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, "not found")
			return
		}
		c.Status(http.StatusOK)
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
