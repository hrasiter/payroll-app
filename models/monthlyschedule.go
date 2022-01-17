package models

type MonthlySchedule struct {
}

func (ms *MonthlySchedule) IsPayDay() bool {
	return true
}
