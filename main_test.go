package main

import (
	"github.com/claytoncasey01/go-oop-to-dod/dod"
	"github.com/claytoncasey01/go-oop-to-dod/oop"
	"testing"
)

var (
	surveys            = oop.MakeSurveys(1)
	surveyInstances    = oop.MakeSurveyInstances(surveys, 1000000)
	scheduled          = make([]oop.SurveyInstance, len(surveyInstances))
	dodSurveys         = dod.MakeSurveys(1, 0)
	dodSurveyInstances = dod.MakeSurveyInstances(1000000, dodSurveys.IDs)
)

func BenchmarkOopMakeSurveysAndInstances(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oop.MakeSurveyInstances(oop.MakeSurveys(1000), 100)
	}
}

func BenchmarkDodMakeSurveysAndInstances(b *testing.B) {
	var newDodSurveys *dod.Surveys
	for i := 0; i < b.N; i++ {
		newDodSurveys = dod.MakeSurveys(1000, 0)
		dod.MakeSurveyInstances(100, newDodSurveys.IDs)
	}
}

func BenchmarkOopSendOrSchedule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oop.SendOrSchedule(surveyInstances, scheduled)
	}
}

func BenchmarkDodSendOrSchedule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dod.SendOrSchedule(&dodSurveyInstances.Sent, dodSurveyInstances.SendDates, &dodSurveyInstances.Scheduled)
	}
}

func BenchmarkOopSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oop.Send(surveyInstances)
	}
}

func BenchmarkDodSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dod.Send(&dodSurveyInstances.Sent)
	}
}
