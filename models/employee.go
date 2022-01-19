package models

type Employee struct {
	EmployeeId int
	Name       string
	PaymentCla PayClassification
	Schedule   PaymentSchedule
	Method     PaymentMethod
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetClassification() PayClassification {
	return e.PaymentCla
}

func (e *Employee) GetSchedule() PaymentSchedule {
	return e.Schedule
}

func (e *Employee) GetMethod() PaymentMethod {
	return e.Method
}

func (e *Employee) GetEmployeeId() int {
	return e.EmployeeId
}
