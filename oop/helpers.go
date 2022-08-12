package oop

import (
	"fmt"
	"time"
)

// MakeSurveys Helper function to generate some random survey data
func MakeSurveys(amount int) []Survey {
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

// MakeSurveyInstances Helper function to generate some random survey instances from some surveys
func MakeSurveyInstances(surveys []Survey, amount int) []SurveyInstance {
	surveyInstances := make([]SurveyInstance, amount*len(surveys))

	for _, survey := range surveys {
		for i := 0; i < amount; i++ {
			sendDate := time.Now().Add(time.Hour * 72)
			surveyInstances[i] = *NewSurveyInstance(i, &survey, sendDate)
		}
	}

	return surveyInstances
}
