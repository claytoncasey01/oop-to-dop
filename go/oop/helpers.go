package oop

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// MakeAuthors Helper function to generate some random author data
func MakeAuthors(amount int) []Author {
	authors := make([]Author, amount)
	for i := 0; i < amount; i++ {
		authors[i] = Author{
			Id:   uuid.New(),
			Name: fmt.Sprintf("Author %d", i),
			Bio:  fmt.Sprintf("I am the bio for Author %d", i),
		}
	}

	return authors
}

// MakePosts Helper function to generate some random Post data
func MakePosts(amount int, authors []Author) []Post {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	p := make([]Post, amount)

	for i := 0; i < amount; i++ {
		randomAuthor := random.Intn(len(authors))
		pTitle := fmt.Sprintf("Post %d", i)
		pBody := fmt.Sprintf("I am the description for survey %d", i)
		pAuthor := authors[randomAuthor]
		p[i] = *NewPost(pAuthor, pTitle, pBody)
	}

	return p
}
