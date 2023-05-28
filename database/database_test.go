package database

import (
	"log"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	db, err := EstablishConnection()
	if err != nil {
		log.Fatalln(err)
	}

	post := NewPost{
		Title:            "Title",
		Content:          "Content",
		PublishTimestamp: time.Now(),
	}

	err = AddPost(db, post)
	if err != nil {
		log.Fatalln(err)
	}

	posts, err := GetAllPosts(db)
	if err != nil {
		log.Fatalln(err)
	}

	postFromDB := posts[len(posts)-1]
	if post.Title != postFromDB.Title ||
		post.Content != postFromDB.Content {
		t.Errorf("FAILED (not equal):\n%#v\n\n%#v", post, postFromDB)
	} else {
		t.Log("SUCCEDED")
	}
}
