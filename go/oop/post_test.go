package oop

import (
	"math/rand"
	"testing"
	"time"
)

const (
	AuthorsAmount            = 100
	PostsAmount              = 10000
	PostsDeterministicAmount = 1000
	PostsPerAuthor           = 10
)

var (
	authors            = MakeAuthors(AuthorsAmount)
	posts              = MakePosts(PostsAmount, authors)
	postsDeterministic = MakePostsDeterministic(PostsDeterministicAmount, PostsPerAuthor, authors)
)

// Unit Tests
func TestOopFindPostById(t *testing.T) {
	uuid := posts[0].Id
	expected := posts[0]
	actual := FindPostById(uuid, posts)

	if actual.Id != expected.Id {
		t.Errorf("Expected %q, Got %q", expected.Id, actual.Id)
	}
}

func TestOopFindPostsByTitle(t *testing.T) {
	expectedTitle := posts[1].Title
	actual := FindPostByTitle("Post 1", posts)

	if actual.Title != expectedTitle {
		t.Errorf("Expected %q, Got %q", expectedTitle, actual.Title)
	}
}

func TestOopFindPostByAuthorName(t *testing.T) {
	var expectedName = authors[0].Name
	var expectedResultLength = PostsPerAuthor
	var actual = FindPostsByAuthorName("Author 0", postsDeterministic)
	var actualName = actual[0].Author.Name
	var actualLength = len(actual)

	if actualLength != expectedResultLength || actualName != expectedName {
		t.Errorf("Expected %q and %d, Got %q and %d", expectedName, expectedResultLength,
			actualName, actualLength)
	}
}

func TestOopPostUpdate(t *testing.T) {
	var updatedPost = PostUpdate{Title: "Post 0 Updated", Body: posts[0].Body + " Updated", Published: false}
	var postToUpdate = postsDeterministic[0]
	postToUpdate.Update(updatedPost)

	if postToUpdate.Title != updatedPost.Title {
		t.Errorf("Expected %q, Got %q", updatedPost.Title, postToUpdate.Title)
	}
}

func TestOopDelete(t *testing.T) {
	var expectedFirstId = postsDeterministic[1].Id
	var expectedLength = len(postsDeterministic) - 1
	var actual = postsDeterministic[0].Delete(postsDeterministic)
	var actualId = actual[0].Id
	var actualLength = len(actual)

	if actualLength != expectedLength || actualId != expectedFirstId {
		t.Errorf("Expected %q and %d, Got %q and %d", expectedFirstId, expectedLength, actualId, actualLength)
	}
}

func TestOopPublish(t *testing.T) {
	var postToPublish = postsDeterministic[1]
	postToPublish.Publish()

	if postToPublish.Published != true {
		t.Errorf("Expected %t, Got %t", true, false)
	}
}

// Benchmarks
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

func BenchmarkOopAddPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		posts = append(posts, *NewPost(authors[0], "Oop Bench Add", "Benchmark add post for Oop"))
	}
}
