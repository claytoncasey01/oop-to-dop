package oop

type QuestionType int

const (
	SingleChoice QuestionType = iota
	MultipleChoice
	Select
	Text
)

func (qt QuestionType) String() string {
	switch qt {
	case SingleChoice:
		return "SingleChoice"
	case MultipleChoice:
		return "MultipleChoice"
	case Select:
		return "Select"
	case Text:
		return "Text"
	}
	return "unknown"
}

type Question struct {
	ID   int
	Name string
	Text string
	Type QuestionType
}

func NewQuestion(id int, questionType QuestionType, name, text string) *Question {
	return &Question{
		ID:   id,
		Name: name,
		Text: text,
		Type: questionType,
	}
}

// FindByType finds a question in the given slice by the given type. If no match, returns nil.
func FindByType(questionType QuestionType, questions []Question) *Question {
	for _, q := range questions {
		if q.Type == questionType {
			return &q
		}
	}

	return nil
}
