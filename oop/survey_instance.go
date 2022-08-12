package oop

import (
	"time"
)

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

// FindBySurveyId finds all survey instances with the matching survey id
func FindBySurveyId(id int, surveyInstances []SurveyInstance) []SurveyInstance {
	instances := make([]SurveyInstance, 10)

	for _, s := range surveyInstances {
		if id == s.SurveyId {
			instances = append(instances, s)
		}
	}

	return instances
}

func SendOrSchedule(instances []SurveyInstance, scheduled []SurveyInstance) []SurveyInstance {
	for _, instance := range instances {
		// Check if we are past the send date or right at
		if instance.SendDate.Before(time.Now()) || instance.SendDate.Equal(time.Now()) {
			instance.Sent = true
		} else {
			scheduled = append(scheduled, instance)
		}
	}

	return scheduled
}

func Send(instances []SurveyInstance) {
	for _, instance := range instances {
		instance.Sent = true
	}
}
