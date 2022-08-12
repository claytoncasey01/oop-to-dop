package dod

import "time"

type SurveyInstances struct {
	IDs       []int
	SurveyIds []int
	Sent      []bool
	SendDates []time.Time
	Scheduled []int
}

func (s *SurveyInstances) Add(SurveyId int) {
	s.IDs = append(s.IDs, len(s.IDs))
	s.SurveyIds = append(s.SurveyIds, SurveyId)
	s.Sent = append(s.Sent, false)
	s.SendDates = append(s.SendDates, time.Now().Add(time.Hour))
}

func SendOrSchedule(sent *[]bool, sendDates []time.Time, scheduled *[]int) {
	for i := range sendDates {
		date := &sendDates[i]
		if date.Before(time.Now()) || date.Equal(time.Now()) {
			(*sent)[i] = true
		} else {
			(*scheduled)[i] = i
		}
	}
}

func Send(sent *[]bool) {
	for i := range *sent {
		(*sent)[i] = true
	}
}

func (s *SurveyInstances) FindBySurveyId(id int) []int {
	found := make([]int, 10)

	for i := range s.SurveyIds {
		if s.SurveyIds[i] == id {
			found = append(found, i)
		}
	}

	return found
}
