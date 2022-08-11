package dod

type QuestionType int

const (
	SingleChoice QuestionType = iota
	MultipleChoice
	Select
	Text
)

type Questions struct {
	IDs   []int
	Names []string
	Text  []string
	Types []QuestionType
}

func (q *Questions) Add(questionType QuestionType, name, text string) {
	q.IDs = append(q.IDs, len(q.IDs))
	q.Names = append(q.Names, name)
	q.Text = append(q.Text, text)
	q.Types = append(q.Types, questionType)
}
