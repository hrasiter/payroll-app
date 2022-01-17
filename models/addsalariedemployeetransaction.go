package models

type AddSalarieEmployeeTransaction struct {
	*AddEmployeeTransaction
	salary float32
}

func NewAddSalariedEmployeeTransaction(id int, name string, salary float32) Transaction {
	aset := &AddSalarieEmployeeTransaction{salary: salary}

	t := &AddEmployeeTransaction{
		aset,
		id,
		name,
	}

	aset.AddEmployeeTransaction = t

	return aset
}

func (t *AddSalarieEmployeeTransaction) GetClassification() PayClassification {
	return &SalariedClassification{salary: t.salary}
}

func (t *AddSalarieEmployeeTransaction) GetSchedule() PaymentSchedule {
	return &MonthlySchedule{}
}
