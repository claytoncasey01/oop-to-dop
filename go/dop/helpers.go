package dop

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// MakeAuthors Helper function to generate some random authors data
func MakeAuthors(amount int) *Authors {
	authors := &Authors{Ids: make([]uuid.UUID, amount), Names: make([]string, amount), Bios: make([]string, amount)}
	for i := 0; i < amount; i++ {
		authors.Ids[i] = uuid.New()
		authors.Names[i] = fmt.Sprintf("Author %d", i)
		authors.Bios[i] = fmt.Sprintf("I am the bio for Author %d", i)
	}

	return authors
}

// MakePosts Helper function to generate some random posts data
func MakePosts(amount int, authors *Authors) *Posts {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	posts := &Posts{Ids: make([]uuid.UUID, amount), Titles: make([]string, amount), Bodies: make([]string, amount),
		UpdateAts: make([]time.Time, amount),
		Published: make([]bool, amount),
		Authors:   make([]int, amount)}
	for i := 0; i < amount; i++ {
		randomAuthor := random.Intn(len(authors.Ids))
		posts.Ids[i] = uuid.New()
		posts.Titles[i] = fmt.Sprintf("Post %d", i)
		posts.Bodies[i] = fmt.Sprintf("I am the body for post %d", i)
		//posts.Published[i] = random.Intn(100)%2 == 0
		posts.Published[i] = false
		posts.Authors[i] = randomAuthor - 1
	}

	return posts
}
