package dop

import (
	"github.com/google/uuid"
	"time"
)

type Posts struct {
	Ids       []uuid.UUID
	Titles    []string
	Bodies    []string
	UpdateAts []time.Time
	Published []bool
	Authors   []int
}

func (p *Posts) Add(title, body string, authorId uuid.UUID, authors *Authors) {
	p.Ids = append(p.Ids, uuid.New())
	p.Titles = append(p.Titles, title)
	p.Bodies = append(p.Bodies, body)
	p.Published = append(p.Published, false)
	p.Authors = append(p.Authors, authors.FindById(authorId))
	p.UpdateAts = append(p.UpdateAts, time.Now())
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

type FindPostsByIdsInput struct {
	PostIds []uuid.UUID
	Ids     []uuid.UUID
}

// FindPostsByIds takes a slice of ids and posts, then finds all posts by the matching ids or an empty slice
func FindPostsByIds(input FindPostsByIdsInput) []int {
	var posts []int

	for i, id := range input.Ids {
		if input.PostIds[i] == id {
			posts = append(posts, i)
		}
	}

	return posts
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

func (p *Posts) Update() {
	// TODO: Implement Update
}

type PublishPostInput struct {
	Ids           []uuid.UUID
	PostIds       []uuid.UUID
	PostPublished []bool
	PostUpdatedAt []time.Time
}

// PublishPost Publishes posts for the given uuids
// TODO: May need to modify this to return a set of new arrays or do an update
func PublishPost(input PublishPostInput) {
	idxs := FindPostsByIds(FindPostsByIdsInput{Ids: input.Ids, PostIds: input.PostIds})

	for _, idx := range idxs {
		if input.PostPublished[idx] {
			continue
		}
		input.PostUpdatedAt[idx] = time.Now()
		input.PostPublished[idx] = true
	}
}

type UpdatePostsInput struct {
	PostIdsToUpdate []uuid.UUID
	Titles          []string
	Bodies          []string
	Posts           *Posts
}

// UpdatePosts Updates any posts with matching ids with the given data.
func UpdatePosts(input UpdatePostsInput) {
	idxs := FindPostsByIds(FindPostsByIdsInput{Ids: input.PostIdsToUpdate, PostIds: input.PostIdsToUpdate})
	for i, idx := range idxs {
		if input.Titles != nil {
			input.Posts.Titles[idx] = input.Titles[i]
		}
		if input.Bodies != nil {
			input.Posts.Bodies[idx] = input.Bodies[i]
		}

		input.Posts.UpdateAts[idx] = time.Now()
	}
}
