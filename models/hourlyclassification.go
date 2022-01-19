package models

type HourlyClassification struct {
	hourlyRate float32
	timecards  map[string]TimeCard
}

func NewHourlyClassification(rate float32) PayClassification {
	return &HourlyClassification{
		hourlyRate: rate,
		timecards:  make(map[string]TimeCard),
	}
}

func (c *HourlyClassification) CalculatePay() float32 {
	return 0
}

func (c *HourlyClassification) GetHourlyRate() float32 {
	return c.hourlyRate
}

func (c *HourlyClassification) GetTimeCard(date string) TimeCard {
	return c.timecards[date]
}

func (hc *HourlyClassification) AddtimeCard(tc TimeCard) {
	hc.timecards[tc.date] = tc
}
