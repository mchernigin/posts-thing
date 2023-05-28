package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mchernigin/posts-thing/database"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

func PostsHandler(db *sqlx.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        posts, err := database.GetAllPosts(db)
        if err != nil {
            log.Fatalln(err)
        }
        c.JSON(http.StatusOK, posts)
    }

    return fn
}

func AuthorsHandler(db *sqlx.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        authors, err := database.GetAllAuthors(db)
        if err != nil {
            log.Fatalln(err)
        }
        c.JSON(http.StatusOK, authors)
    }

    return fn
}

func setupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/posts", PostsHandler(db))
	r.GET("/authors", AuthorsHandler(db))

    return r
}

func main() {
	db, err := database.EstablishConnection()
	if err != nil {
		log.Fatalln(err)
	}

    // gin.SetMode(gin.ReleaseMode)
    r := setupRouter(db)
	r.Run()

	// post := database.NewPost{
	// 	Title:            "Title",
	// 	Content:          "Content",
	// 	PublishTimestamp: time.Now(),
	// }
	// err = database.AddPost(db, post)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// author := database.NewAuthor{
	// 	Name:    "Michael",
	// 	Surname: "Chernigin",
	// 	Website: "https://chernigin.com",
	// 	Status:  "aboba",
	// }
	// err = database.AddAuthor(db, author)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
