package oop

import "github.com/google/uuid"

type Author struct {
	Id   uuid.UUID
	Name string
	Bio  string
}

func NewAuthor(name, bio string) *Author {
	return &Author{
		Id:   uuid.New(),
		Name: name,
		Bio:  bio,
	}
}

// FindAuthorById Takes a slice of authors and returns a pointer to an Author if found or nil if not.
func FindAuthorById(id uuid.UUID, authors []Author) *Author {
	for _, a := range authors {
		if a.Id == id {
			return &a
		}
	}
	return nil
}
