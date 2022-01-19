package models

type TimeCard struct {
	hours int
	date  string
}

func (tc *TimeCard) GetHours() int {
	return tc.hours
}
