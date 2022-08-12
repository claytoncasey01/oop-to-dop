package main

import (
	"fmt"
	"github.com/claytoncasey01/go-oop-to-dod/dod"
	"github.com/claytoncasey01/go-oop-to-dod/oop"
)

func main() {
	surveys := oop.MakeSurveys(10)
	surveyInstances := oop.MakeSurveyInstances(surveys, 1000)
	scheduled := make([]oop.SurveyInstance, len(surveyInstances))
	fmt.Printf("SurveyInstances Length: %d\n", len(surveyInstances))
	// Send all our surveys
	scheduled = oop.SendOrSchedule(surveyInstances, scheduled)

	dodSurveys := dod.MakeSurveys(10, 0)
	dodSurveyInstances := dod.MakeSurveyInstances(1000, dodSurveys.IDs)
	fmt.Printf("DodSurveys Length: %d\n", len(dodSurveyInstances.IDs))

}
