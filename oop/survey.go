package oop

// Survey represents the basic stucture of a survey
type Survey struct {
	ID          int
	Name        string
	Description string
	Questions   []Question
}

// NewSurvey creates a new survey instance
func NewSurvey(id int, questions []Question, name, description string) *Survey {
	return &Survey{
		ID:          id,
		Name:        name,
		Description: description,
		Questions:   questions,
	}
}
