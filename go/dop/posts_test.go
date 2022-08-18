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

// Unit Tests
func TestDopAddPost(t *testing.T) {
	expectedLength := len(posts.Ids) + 1
	// Add the actual post
	AddPost("New TestDop Post", "TestDop Post Body", authors.Ids[0], posts, authors)
	actualLength := len(posts.Ids)

	if actualLength != expectedLength {
		t.Errorf("Expected %d posts, got %d", expectedLength, actualLength)
	}
}

func TestDopFindPostById(t *testing.T) {
	postId := posts.Ids[0]
	expected := 0
	actual := FindPostById(FindPostByIdInput{
		Ids: posts.Ids,
		Id:  postId,
	})

	if actual != expected {
		t.Errorf("Expected %d posts, got %d", expected, actual)
	}

}

func TestDopUpdatePost(t *testing.T) {
	postId := posts.Ids[0]
	expectedTitle := posts.Titles[0] + " Updated"
	expectedBody := posts.Bodies[0] + " Updated"
	UpdatePost(UpdatePostInput{PostIdToUpdate: postId, Title: expectedTitle, Body: expectedBody, Posts: posts})
	actualTitle := posts.Titles[0]
	actualBody := posts.Bodies[0]

	if actualTitle != expectedTitle || actualBody != expectedBody {
		t.Errorf("Expected %s and %s, got %s and %s", expectedTitle, expectedBody, actualTitle, actualBody)
	}
}

func TestDopPublishPost(t *testing.T) {
	postId := posts.Ids[0]
	PublishPost(PublishPostInput{Id: postId, PostIds: posts.Ids, PostPublished: posts.Published, PostUpdatedAt: posts.UpdatedAts})
	actual := posts.Published[0]

	if actual != true {
		t.Errorf("Expected %t, got %t", true, actual)
	}
}

func TestDopDeletePost(t *testing.T) {
	postId := posts.Ids[0]
	expectedLength := len(posts.Ids) - 1
	DeletePost(DeletePostInput{DeleteId: postId, Posts: posts})
	actualLength := len(posts.Ids)

	if actualLength != expectedLength {
		t.Errorf("Expected %d posts, got %d", expectedLength, actualLength)
	}
}

// Benchmarks
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
		Id:            posts.Ids[0], // This normally wouldn't be the same thing but lazy, so just publishing all posts
		PostIds:       posts.Ids,
		PostUpdatedAt: posts.UpdatedAts,
		PostPublished: posts.Published,
	}
	for i := 0; i < b.N; i++ {
		PublishPost(input)
	}
}

func BenchmarkDopUpdatePosts(b *testing.B) {
	updatedTitle := posts.Titles[0] + " Updated"
	updatedBody := posts.Bodies[0] + " Updated"

	input := UpdatePostInput{PostIdToUpdate: posts.Ids[0], Title: updatedTitle, Body: updatedBody, Posts: posts}

	for i := 0; i < b.N; i++ {
		UpdatePost(input)
	}
}

func BenchmarkDopAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddPost("TestDop Add", "Post to bench adding", authors.Ids[0], posts, authors)
	}
}

func BenchmarkDopDelete(b *testing.B) {
	input := DeletePostInput{}

	for i := 0; i < b.N; i++ {
		if len(posts.Ids) > 0 {
			input.Posts = posts
			input.DeleteId = posts.Ids[0]
			DeletePost(input)
		}
	}
}
