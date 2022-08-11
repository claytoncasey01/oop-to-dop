package main

import (
	"fmt"
	"time"
)

func main() {
	surveys := makeSurveys(10)
	surveyInstances := makeSurveyInstances(surveys, 1000)

	// Send all our surveys
	for _, surveyInstance := range surveyInstances {
		surveyInstance.Send()
		fmt.Printf("Survey Instance %d for %s has a sent status of %t\n", surveyInstance.ID, surveys[surveyInstance.SurveyId].Name, surveyInstance.Sent)
	}

}

// Helper function to generate some random survey data
func makeSurveys(amount int) []Survey {
	s := make([]Survey, amount)

	for i := 0; i < amount; i++ {
		qName := fmt.Sprintf("Question %d", i)
		qText := fmt.Sprintf("I am the text for question %d", i)
		q := NewQuestion(i, SingleChoice, qName, qText)
		sName := fmt.Sprintf("Survey %d", i)
		sDesc := fmt.Sprintf("I am the description for survey %d", i)
		s[i] = *NewSurvey(i, []Question{*q}, sName, sDesc)
	}

	return s
}

func makeSurveyInstances(surveys []Survey, amount int) []SurveyInstance {
	surveyInstances := make([]SurveyInstance, amount)

	for _, survey := range surveys {
		for i := 0; i < amount; i++ {
			sendDate := time.Now().Add(time.Hour * 72)
			surveyInstances[i] = *NewSurveyInstance(i, &survey, sendDate)
		}
	}

	return surveyInstances
}
