package models

var PayrollDB = PayrollDatabase{make(map[int]*Employee)}

type PayrollDatabase struct {
	database map[int]*Employee
}

func (db *PayrollDatabase) GetEmployee(id int) *Employee {
	return db.database[id]
}

func (db *PayrollDatabase) AddEmployee(id int, e *Employee) {
	db.database[id] = e
}
