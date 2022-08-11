package main

import (
	"github.com/claytoncasey01/go-oop-to-dod/dod"
	"github.com/claytoncasey01/go-oop-to-dod/oop"
	"testing"
)

var (
	surveys            = makeOopSurveys(1)
	surveyInstances    = makeOopsSurveyInstances(surveys, 1000000)
	scheduled          = make([]oop.SurveyInstance, len(surveyInstances))
	dodSurveys         = makeDodSurveys(1, 0)
	dodSurveyInstances = makeDodSurveyInstances(1000000, dodSurveys.IDs)
)

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
