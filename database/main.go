package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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

func GetPostById(db *sqlx.DB, id int64) (Post, error) {
	post := Post{}
	err := db.Get(&post, "select * from posts where id = $1", id)
	return post, err
}

func DeletePostById(db *sqlx.DB, id int64) error {
	_, err := db.Exec("delete from posts where id = $1", id)
	return err
}

func AddPost(db *sqlx.DB, post NewPost) (int64, error) {
	var id int64
	err := db.QueryRowx( // NOTE: Sadly there is no Named func for this
		`insert into posts (title, content, publish_timestamp)
         values ($1, $2, $3) returning id`,
		post.Title, post.Content, post.PublishTimestamp).Scan(&id)

	return id, err
}

func GetAllAuthors(db *sqlx.DB) ([]Author, error) {
	authors := []Author{}
	err := db.Select(&authors, "select * from authors")
	return authors, err
}

func AddAuthor(db *sqlx.DB, author NewAuthor) (int64, error) {
	var id int64
	err := db.QueryRowx( // NOTE: Sadly there is no Named func for this
		`insert into authors (name, surname, website, status)
         values ($1, $2, $3, $4) returning id`,
		author.Name, author.Surname, author.Website, author.Status).Scan(&id)

	return id, err
}

func GetAuthorById(db *sqlx.DB, id int64) (Author, error) {
	author := Author{}
	err := db.Get(&author, "select * from authors where id = $1", id)
	return author, err
}

func DeleteAuthorById(db *sqlx.DB, id int64) error {
	_, err := db.Exec("delete from authors where id = $1", id)
	return err
}
