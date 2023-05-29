package database

import "time"

type NewPost struct {
	Title            string    `db:"title"`
	Content          string    `db:"content"`
	PublishTimestamp time.Time `db:"publish_timestamp"`
}

type Post struct {
	Id int64 `db:"id"`
	NewPost
}

type NewAuthor struct {
	Name    string `db:"name"`
	Surname string `db:"surname"`
	Status  string `db:"status"`
	Website string `db:"website"`
}

type Author struct {
	Id int64 `db:"id"`
	NewAuthor
}
