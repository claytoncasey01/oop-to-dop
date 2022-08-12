package dod

type Surveys struct {
	IDs          []int
	Names        []string
	Descriptions []string
	QuestionIds  [][]int
}

func (s *Surveys) Add(name, description string, questionIds []int) {
	s.IDs = append(s.IDs, len(s.IDs))
	s.Names = append(s.Names, name)
	s.Descriptions = append(s.Descriptions, description)
	s.QuestionIds = append(s.QuestionIds, questionIds)
}

// FindById returns the index of the survey if found
func (s *Surveys) FindById(id int) int {
	for i := range s.IDs {
		if s.IDs[i] == id {
			return i
		}
	}

	return -1
}

// FindByName returns the index of the survey if found
func (s *Surveys) FindByName(name string) int {
	for i := range s.Names {
		if s.Names[i] == name {
			return i
		}
	}

	return -1
}
