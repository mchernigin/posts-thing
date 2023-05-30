package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mchernigin/posts-thing/server/routes"
)

func Serve(db *sqlx.DB) {
	router := gin.Default()

	routes.SetupPostsRoutes(router, db)
	routes.SetupAuthorsRoutes(router, db)

	router.Run()
}
