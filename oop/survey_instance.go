package oop

import "time"

type SurveyInstance struct {
	ID       int
	SurveyId int
	Sent     bool
	SendDate time.Time
}

func NewSurveyInstance(id int, survey *Survey, sendDate time.Time) *SurveyInstance {
	return &SurveyInstance{
		ID:       id,
		SurveyId: survey.ID,
		Sent:     false,
		SendDate: sendDate,
	}
}

func (s *SurveyInstance) Send() {
	s.Sent = true
}
