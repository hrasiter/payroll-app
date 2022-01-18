package models

type HourlyClassification struct {
	hourlyRate float32
}

func (c *HourlyClassification) CalculatePay() float32 {
	return 0
}

func (c *HourlyClassification) GetHourlyRate() float32 {
	return c.hourlyRate
}
