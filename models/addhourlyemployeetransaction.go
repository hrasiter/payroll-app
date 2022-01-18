package models

type AddHourlyEmployeeTransaction struct {
	hourlyRate float32
	*AddEmployeeTransaction
}

func NewAddHourlyEmployeeTransaction(id int, name string, rate float32) *AddHourlyEmployeeTransaction {
	ahet := &AddHourlyEmployeeTransaction{hourlyRate: rate}

	t := &AddEmployeeTransaction{
		ahet,
		id,
		name,
	}

	ahet.AddEmployeeTransaction = t

	return ahet
}

func (t *AddHourlyEmployeeTransaction) GetClassification() PayClassification {
	return &HourlyClassification{hourlyRate: t.hourlyRate}
}

func (t *AddHourlyEmployeeTransaction) GetSchedule() PaymentSchedule {
	return &WeeklySchedule{}
}
