package database

import (
	"github.com/joho/godotenv"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func EstablishConnection() (*sqlx.DB, error) {
	godotenv.Load()
	db := os.Getenv("DATABASE_URL")
	return sqlx.Connect("postgres", db)
}

func GetAllPosts(db *sqlx.DB) ([]Post, error) {
	posts := []Post{}
	err := db.Select(&posts, "select * from posts")
	return posts, err
}

func AddPost(db *sqlx.DB, post NewPost) error {
	_, err := db.NamedExec(`insert into posts (title, content, publish_timestamp)
                            values (:title, :content, :publish_timestamp)`, post)
	return err
}

func GetAllAuthors(db *sqlx.DB) ([]Author, error) {
	authors := []Author{}
	err := db.Select(&authors, "select * from authors")
	return authors, err
}

func AddAuthor(db *sqlx.DB, author NewAuthor) error {
	_, err := db.NamedExec(`insert into authors (name, surname, website, status)
                            values (:name, :surname, :website, :status)`,
		author)
	return err
}
