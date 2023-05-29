package main

import (
	"github.com/joho/godotenv"
	"github.com/mchernigin/posts-thing/database"
	"github.com/mchernigin/posts-thing/server"
	"log"
)

func main() {
	godotenv.Load()
	db, err := database.EstablishConnection()
	if err != nil {
		log.Fatalln(err)
	}

	server.Serve(db)
}
