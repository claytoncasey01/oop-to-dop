package oop

import (
	"math/rand"
	"testing"
	"time"
)

var (
	authors = MakeAuthors(100)
	posts   = MakePosts(10000, authors)
)

func BenchmarkOopFindPostById(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPost := random.Intn(len(posts) - 1)

	for i := 0; i < b.N; i++ {
		FindPostById(posts[randomPost].Id, posts)
	}
}

func BenchmarkOopFindPostByTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPostByTitle("Post 1", posts)
	}
}

func BenchmarkOopFindPostByAuthorName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPostsByAuthorName("Casey", posts)
	}
}

func BenchmarkOopPublishPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, p := range posts {
			if !p.Published {
				p.Publish()
			}
		}
	}
}

func BenchmarkOopUpdatePost(b *testing.B) {
	for _, p := range posts {
		for i := 0; i < b.N; i++ {
			p.Update(PostUpdate{Title: p.Title + "Updated", Body: p.Body + "Updated"})
		}
	}
}
