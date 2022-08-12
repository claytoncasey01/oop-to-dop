package dod

import (
	"fmt"
	"time"
)

// MakeQuestions Helper function to generate questions and return a slice of questionIds
func MakeQuestions(amount int) []int {
	questions := &Questions{IDs: make([]int, amount), Names: make([]string, amount), Text: make([]string, amount), Types: make([]QuestionType, amount)}
	questionIds := make([]int, amount)

	for i := 0; i < amount; i++ {
		questions.IDs[i] = i
		questions.Names[i] = fmt.Sprintf("Question %d", i)
		questions.Text[i] = fmt.Sprintf("I am the text for question %d", i)
		questions.Types[i] = SingleChoice
		questionIds[i] = i
	}

	return questionIds
}

// MakeSurveys Helper function to generate some random survey data
func MakeSurveys(amount, questionAmount int) *Surveys {
	surveys := &Surveys{IDs: make([]int, amount), Names: make([]string, amount), Descriptions: make([]string, amount), QuestionIds: make([][]int, amount)}
	for i := 0; i < amount; i++ {
		surveys.IDs[i] = i
		surveys.Names[i] = fmt.Sprintf("Survey %d", i)
		surveys.Descriptions[i] = fmt.Sprintf("I am the description for survey %d", i)
		surveys.QuestionIds[i] = MakeQuestions(questionAmount)
	}

	return surveys
}

// MakeSurveyInstances Helper function to generate some random survey instances from some surveys
func MakeSurveyInstances(amount int, surveyIds []int) *SurveyInstances {
	trueAmount := amount * len(surveyIds)
	surveyInstances := &SurveyInstances{IDs: make([]int, trueAmount), SurveyIds: make([]int, trueAmount), Sent: make([]bool, trueAmount), SendDates: make([]time.Time, trueAmount), Scheduled: make([]int, trueAmount)}

	for idx := range surveyIds {
		id := &surveyIds[idx]
		for i := 0; i < amount; i++ {
			surveyInstances.IDs[i] = i
			surveyInstances.SurveyIds[i] = *id
			surveyInstances.Sent[i] = false
			surveyInstances.SendDates[i] = time.Now()
		}
	}

	return surveyInstances
}
