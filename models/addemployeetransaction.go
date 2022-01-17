package models

type AddEmployeeTransaction struct {
	AddEmployee

	empId int
	name  string
}

func (t *AddEmployeeTransaction) Execute() {
	pc := t.GetClassification()
	ps := t.GetSchedule()
	pm := &HoldMethod{}
	e := &Employee{
		EmployeeId: t.empId,
		Name:       t.name,
		PaymentCla: pc,
		Schedule:   ps,
		Method:     pm,
	}
	PayrollDB.AddEmployee(t.empId, e)
}
