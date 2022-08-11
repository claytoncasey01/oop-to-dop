package main

import (
	"fmt"
	"github.com/claytoncasey01/go-oop-to-dod/dod"
	"github.com/claytoncasey01/go-oop-to-dod/oop"
	"time"
)

func main() {
	surveys := makeOopSurveys(10)
	surveyInstances := makeOopsSurveyInstances(surveys, 1000)
	scheduled := make([]oop.SurveyInstance, len(surveyInstances))
	fmt.Printf("SurveyInstances Length: %d\n", len(surveyInstances))
	// Send all our surveys
	scheduled = oop.SendOrSchedule(surveyInstances, scheduled)

	dodSurveys := makeDodSurveys(10, 0)
	dodSurveyInstances := makeDodSurveyInstances(1000, dodSurveys.IDs)
	fmt.Printf("DodSurveys Length: %d\n", len(dodSurveyInstances.IDs))

}

// Helper function to generate some random survey data
func makeOopSurveys(amount int) []oop.Survey {
	s := make([]oop.Survey, amount)

	for i := 0; i < amount; i++ {
		qName := fmt.Sprintf("Question %d", i)
		qText := fmt.Sprintf("I am the text for question %d", i)
		q := oop.NewQuestion(i, oop.SingleChoice, qName, qText)
		sName := fmt.Sprintf("Survey %d", i)
		sDesc := fmt.Sprintf("I am the description for survey %d", i)
		s[i] = *oop.NewSurvey(i, []oop.Question{*q}, sName, sDesc)
	}

	return s
}

func makeOopsSurveyInstances(surveys []oop.Survey, amount int) []oop.SurveyInstance {
	surveyInstances := make([]oop.SurveyInstance, amount*len(surveys))

	for _, survey := range surveys {
		for i := 0; i < amount; i++ {
			sendDate := time.Now().Add(time.Hour * 72)
			surveyInstances[i] = *oop.NewSurveyInstance(i, &survey, sendDate)
		}
	}

	return surveyInstances
}

func makeDodQuestions(amount int) []int {
	questions := &dod.Questions{IDs: make([]int, amount), Names: make([]string, amount), Text: make([]string, amount), Types: make([]dod.QuestionType, amount)}
	questionIds := make([]int, amount)

	for i := 0; i < amount; i++ {
		questions.IDs[i] = i
		questions.Names[i] = fmt.Sprintf("Question %d", i)
		questions.Text[i] = fmt.Sprintf("I am the text for question %d", i)
		questions.Types[i] = dod.SingleChoice
		questionIds[i] = i
	}

	return questionIds
}

func makeDodSurveys(amount, questionAmount int) *dod.Surveys {
	dodSurveys := &dod.Surveys{IDs: make([]int, amount), Names: make([]string, amount), Descriptions: make([]string, amount), QuestionIds: make([][]int, amount)}
	for i := 0; i < amount; i++ {
		dodSurveys.IDs[i] = i
		dodSurveys.Names[i] = fmt.Sprintf("Survey %d", i)
		dodSurveys.Descriptions[i] = fmt.Sprintf("I am the description for survey %d", i)
		dodSurveys.QuestionIds[i] = makeDodQuestions(questionAmount)
	}

	return dodSurveys
}

func makeDodSurveyInstances(amount int, surveyIds []int) *dod.SurveyInstances {
	trueAmount := amount * len(surveyIds)
	dodSurveyInstances := &dod.SurveyInstances{IDs: make([]int, trueAmount), SurveyIds: make([]int, trueAmount), Sent: make([]bool, trueAmount), SendDates: make([]time.Time, trueAmount), Scheduled: make([]int, trueAmount)}

	for idx := range surveyIds {
		id := &surveyIds[idx]
		for i := 0; i < amount; i++ {
			dodSurveyInstances.IDs[i] = i
			dodSurveyInstances.SurveyIds[i] = *id
			dodSurveyInstances.Sent[i] = false
			dodSurveyInstances.SendDates[i] = time.Now()
		}
	}

	return dodSurveyInstances
}
