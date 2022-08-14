package dod

import "github.com/google/uuid"

type Authors struct {
	Ids   []uuid.UUID
	Names []string
	Bios  []string
}

func (a *Authors) Add(name, bio string) {
	a.Ids = append(a.Ids, uuid.New())
	a.Names = append(a.Names, name)
	a.Bios = append(a.Bios, bio)
}

func (a *Authors) FindById(id uuid.UUID) int {
	for i := range a.Ids {
		if a.Ids[i] == id {
			return i
		}
	}

	return -1
}

func (a *Authors) FindByName(name string) int {
	for i := range a.Names {
		if a.Names[i] == name {
			return i
		}
	}
	return -1
}

type AuthorsByName struct {
	name  string
	names []string
}

func FindAuthorByName(findBy AuthorsByName) int {
	for i := range findBy.names {
		if findBy.names[i] == findBy.name {
			return i
		}
	}

	return -1
}
