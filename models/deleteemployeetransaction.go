package models

type DeleteEmployeeTransaction struct {
	empId int
}

func NewDeleteEmployeeTransaction(empId int) Transaction {
	return &DeleteEmployeeTransaction{empId}
}

func (t *DeleteEmployeeTransaction) Execute() {
	PayrollDB.DeleteEmployee(t.empId)
}
