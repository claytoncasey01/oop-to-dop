package dop

import (
	"math/rand"
	"testing"
	"time"
)

var (
	authors = MakeAuthors(100)
	posts   = MakePosts(10000, authors)
)

func BenchmarkDodFindById(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPost := random.Intn(len(posts.Ids) - 1)

	for i := 0; i < b.N; i++ {
		posts.FindById(posts.Ids[randomPost])
	}
}

func BenchmarkDodFindByTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		posts.FindByTitle("Post 1")
	}
}

func BenchmarkDodFindByAuthorName(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomAuthor := random.Intn(len(authors.Ids) - 1)

	for i := 0; i < b.N; i++ {
		posts.FindByAuthorName(authors.Names[randomAuthor], authors)
	}
}

func BenchmarkDodFindPostsByAuthorName(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomAuthorIdx := random.Intn(len(authors.Ids) - 1)
	randomAuthor := authors.Names[randomAuthorIdx]

	for i := 0; i < b.N; i++ {
		FindPostsByAuthorName(PostsByAuthorName{name: randomAuthor, authorNames: authors.Names})
	}
}

func BenchmarkDodPublishPosts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		posts.Publish(posts.Ids)
	}
}
