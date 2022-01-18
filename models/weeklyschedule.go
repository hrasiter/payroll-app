package models

type WeeklySchedule struct {
}

func (s *WeeklySchedule) IsPayDay() bool {
	return true
}
