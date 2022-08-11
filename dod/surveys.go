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
