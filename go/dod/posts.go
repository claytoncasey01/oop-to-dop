package dod

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

// FindById returns the index of the survey if found
func (p *Posts) FindById(id uuid.UUID) int {
	for i := range p.Ids {
		if p.Ids[i] == id {
			return i
		}
	}

	return -1
}

func (p *Posts) FindByIds(ids []uuid.UUID) []int {
	var posts []int

	for i, id := range ids {
		if p.Ids[i] == id {
			posts = append(posts, i)
		}
	}

	return posts
}

// FindByTitle returns the index of the survey if found
func (p *Posts) FindByTitle(title string) int {
	for i := range p.Titles {
		if p.Titles[i] == title {
			return i
		}
	}

	return -1
}

// FindByAuthorName returns the indexes of a posts by the given author or empty slice if none are found
// TODO: Improve this
func (p *Posts) FindByAuthorName(name string, authors *Authors) []int {
	authorIdx := authors.FindByName(name)
	var posts []int

	for i, author := range p.Authors {
		if author == authorIdx {
			posts = append(posts, i)
		}
	}

	return posts
}

type PostsByAuthorName struct {
	name        string
	authorNames []string
}

func FindPostsByAuthorName(input PostsByAuthorName) []int {
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

// Publish Publishes posts for the given uuids
func (p *Posts) Publish(ids []uuid.UUID) {
	idxs := p.FindByIds(ids)

	for _, idx := range idxs {
		if p.Published[idx] {
			continue
		}
		p.UpdateAts[idx] = time.Now()
		p.Published[idx] = true
	}
}
