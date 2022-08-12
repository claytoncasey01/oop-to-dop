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

// FindById Takes a slice of surveys and an id, and returns the survey or nil if not found
func FindById(id int, surveys []Survey) *Survey {
	for _, s := range surveys {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

// FindByName Takes a slice of survey and a name, and returns the survey or nil if not found
func FindByName(name string, surveys []Survey) *Survey {
	for _, s := range surveys {
		if s.Name == name {
			return &s
		}
	}
	return nil
}
