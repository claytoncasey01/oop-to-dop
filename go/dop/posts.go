package dop

import (
	"github.com/google/uuid"
	"time"
)

type Posts struct {
	Ids        []uuid.UUID
	Titles     []string
	Bodies     []string
	UpdatedAts []time.Time
	Published  []bool
	Authors    []int
}

func AddPost(title, body string, authorId uuid.UUID, p *Posts, authors *Authors) {
	p.Ids = append(p.Ids, uuid.New())
	p.Titles = append(p.Titles, title)
	p.Bodies = append(p.Bodies, body)
	p.Published = append(p.Published, false)
	p.Authors = append(p.Authors, authors.FindById(authorId))
	p.UpdatedAts = append(p.UpdatedAts, time.Now())
}

type FindPostByIdInput struct {
	Ids []uuid.UUID
	Id  uuid.UUID
}

// FindPostById returns the index of the survey if found
func FindPostById(input FindPostByIdInput) int {
	for i := range input.Ids {
		if input.Ids[i] == input.Id {
			return i
		}
	}

	return -1
}

type FindPostByTitleInput struct {
	PostTitles []string
	Title      string
}

// FindPostByTitle returns the index of the survey if found
func FindPostByTitle(input FindPostByTitleInput) int {
	for i := range input.PostTitles {
		if input.PostTitles[i] == input.Title {
			return i
		}
	}

	return -1
}

type PostsByAuthorNameInput struct {
	name        string
	authorNames []string
}

func FindPostsByAuthorName(input PostsByAuthorNameInput) []int {
	var posts []int

	for i, name := range input.authorNames {
		if name == input.name {
			posts = append(posts, i)
		}
	}

	return posts
}

type PublishPostInput struct {
	Id            uuid.UUID
	PostIds       []uuid.UUID
	PostPublished []bool
	PostUpdatedAt []time.Time
}

// PublishPost Publishes posts for the given uuids
// TODO: May need to modify this to return a set of new arrays or do an update
func PublishPost(input PublishPostInput) {
	var idx = FindPostById(FindPostByIdInput{Ids: input.PostIds, Id: input.Id})

	if input.PostPublished[idx] {
		return
	}
	input.PostUpdatedAt[idx] = time.Now()
	input.PostPublished[idx] = true
}

type UpdatePostInput struct {
	PostIdToUpdate uuid.UUID
	Title          string
	Body           string
	Posts          *Posts
}

// UpdatePost Updates any posts with matching ids with the given data.
func UpdatePost(input UpdatePostInput) {
	var idx = FindPostById(FindPostByIdInput{Ids: input.Posts.Ids, Id: input.PostIdToUpdate})

	if input.Title != "" {
		input.Posts.Titles[idx] = input.Title
	}
	if input.Body != "" {
		input.Posts.Bodies[idx] = input.Body
	}

	input.Posts.UpdatedAts[idx] = time.Now()
}

type DeletePostInput struct {
	DeleteId uuid.UUID
	Posts    *Posts
}

func DeletePost(input DeletePostInput) {
	var idx = FindPostById(FindPostByIdInput{Ids: input.Posts.Ids, Id: input.DeleteId})
	input.Posts.Ids = append(input.Posts.Ids[:idx], input.Posts.Ids[idx+1:]...)
	input.Posts.Titles = append(input.Posts.Titles[:idx], input.Posts.Titles[idx+1:]...)
	input.Posts.Bodies = append(input.Posts.Bodies[:idx], input.Posts.Bodies[idx+1:]...)
	input.Posts.Published = append(input.Posts.Published[:idx], input.Posts.Published[idx+1:]...)
	input.Posts.UpdatedAts = append(input.Posts.UpdatedAts[:idx], input.Posts.UpdatedAts[idx+1:]...)
	input.Posts.Authors = append(input.Posts.Authors[:idx], input.Posts.Authors[idx+1:]...)
}
