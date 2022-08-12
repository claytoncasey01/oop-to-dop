package dod

import "testing"

var (
	surveys         = MakeSurveys(1, 0)
	surveyInstances = MakeSurveyInstances(1000000, surveys.IDs)
)

func BenchmarkDodFindById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		surveys.FindById(1)
	}
}

func BenchmarkDodFindByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		surveys.FindByName("Survey 1")
	}
}

func BenchmarkDodFindBySurveyId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		surveyInstances.FindBySurveyId(1)
	}
}

func BenchmarkDodMakeSurveysAndInstances(b *testing.B) {
	var newDodSurveys *Surveys
	for i := 0; i < b.N; i++ {
		newDodSurveys = MakeSurveys(1000, 0)
		MakeSurveyInstances(100, newDodSurveys.IDs)
	}
}

func BenchmarkDodSendOrSchedule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SendOrSchedule(&surveyInstances.Sent, surveyInstances.SendDates, &surveyInstances.Scheduled)
	}
}

func BenchmarkDodSend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Send(&surveyInstances.Sent)
	}
}
