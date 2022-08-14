package oop

import (
	"github.com/google/uuid"
	"time"
)

// Post represents the basic stucture of a survey
type Post struct {
	Id        uuid.UUID
	Title     string
	Body      string
	Published bool
	UpdatedAt time.Time
	Author    Author
}

// NewPost creates a new survey instance
func NewPost(author Author, title, description string) *Post {
	return &Post{
		Id:        uuid.New(),
		Title:     title,
		Published: false,
		UpdatedAt: time.Now(),
		Body:      description,
		Author:    author,
	}
}

type PostUpdate struct {
	Title     string
	Body      string
	Published bool
}

// Update updates a post
func (p *Post) Update(updatedPost PostUpdate) {
	p.UpdatedAt = time.Now()
	p.Body = updatedPost.Body
	p.Published = updatedPost.Published
	p.Title = updatedPost.Title
}

func (p *Post) Publish() {
	p.UpdatedAt = time.Now()
	p.Published = true
}

// FindPostById Takes a slice of posts and an id, and returns the survey or nil if not found
func FindPostById(id uuid.UUID, posts []Post) *Post {
	for _, s := range posts {
		if s.Id == id {
			return &s
		}
	}
	return nil
}

// FindPostByTitle Takes a slice of survey and a name, and returns the survey or nil if not found
func FindPostByTitle(title string, posts []Post) *Post {
	for _, p := range posts {
		if p.Title == title {
			return &p
		}
	}
	return nil
}

// FindPostsByAuthorName Takes a slice of posts and an author name, and returns all matching posts or empty slice if
// none found
func FindPostsByAuthorName(authorName string, posts []Post) []Post {
	var foundPosts []Post

	for _, p := range posts {
		if p.Author.Name == authorName {
			foundPosts = append(foundPosts, p)
		}
	}
	return foundPosts
}
