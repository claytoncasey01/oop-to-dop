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

func BenchmarkDopFindById(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomPost := random.Intn(len(posts.Ids) - 1)
	input := FindPostByIdInput{Ids: posts.Ids, Id: posts.Ids[randomPost]}

	for i := 0; i < b.N; i++ {
		FindPostById(input)
	}
}

func BenchmarkDopFindByTitle(b *testing.B) {
	input := FindPostByTitleInput{Title: "Post 1", PostTitles: posts.Titles}

	for i := 0; i < b.N; i++ {
		FindPostByTitle(input)
	}
}

func BenchmarkDopFindPostsByAuthorName(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomAuthorIdx := random.Intn(len(authors.Ids) - 1)
	randomAuthor := authors.Names[randomAuthorIdx]
	input := PostsByAuthorNameInput{name: randomAuthor, authorNames: authors.Names}

	for i := 0; i < b.N; i++ {
		FindPostsByAuthorName(input)
	}
}

func BenchmarkDopPublishPosts(b *testing.B) {
	input := PublishPostInput{
		Ids:           posts.Ids, // This normally wouldn't be the same thing but lazy, so just publishing all posts
		PostIds:       posts.Ids,
		PostUpdatedAt: posts.UpdateAts,
		PostPublished: posts.Published,
	}
	for i := 0; i < b.N; i++ {
		PublishPost(input)
	}
}

func BenchmarkDopUpdatePosts(b *testing.B) {
	var updatedTitles = make([]string, len(posts.Titles))
	var updatedBodies = make([]string, len(posts.Bodies))

	for i := 0; i < len(updatedTitles); i++ {
		updatedTitles[i] = posts.Titles[i] + " Updated"
		updatedBodies[i] = posts.Bodies[i] + " Updated"
	}
	input := UpdatePostsInput{PostIdsToUpdate: posts.Ids, Titles: updatedTitles, Bodies: updatedBodies, Posts: posts}

	for i := 0; i < b.N; i++ {
		UpdatePosts(input)
	}
}

func BenchmarkDopAdd(b *testing.B) {
  for i := 0; i < b.N; i++ {
    AddPost("Test Add", "Post to bench adding", authors.Ids[0], posts, authors)
  }
}
