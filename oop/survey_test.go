package oop

import "testing"

var (
	surveys         = MakeSurveys(1)
	surveyInstances = MakeSurveyInstances(surveys, 1000000)
	scheduled       = make([]SurveyInstance, len(surveyInstances))
)

func BenchmarkOopFindById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindById(1, surveys)
	}
}

func BenchmarkOopFindByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindByName("Survey 1", surveys)
	}
}

func BenchmarkOopFindBySurveyId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindBySurveyId(1, surveyInstances)
	}
}

func BenchmarkOopMakeSurveysAndInstances(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSurveyInstances(MakeSurveys(1000), 100)
	}
}

func BenchmarkOopSendOrSchedule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendOrSchedule(surveyInstances, scheduled)
	}
}

func BenchmarkOopSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Send(surveyInstances)
	}
}
