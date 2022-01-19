package models

import "testing"

func Test_DeleteEmployee(t *testing.T) {
	t.Run("Delete An Employee from DB", func(t *testing.T) {
		empid := 1
		name := "Rasit"
		salary := float32(5000)

		transaction := NewAddSalariedEmployeeTransaction(empid, name, salary)

		transaction.Execute()

		e := PayrollDB.GetEmployee(empid)

		if e.GetName() != name {
			t.Errorf("Cannot add employee")
		}

		transaction = NewDeleteEmployeeTransaction(empid)

		transaction.Execute()

		delemp := PayrollDB.GetEmployee(empid)

		if delemp != nil {
			t.Errorf("Cannot delete employee")
		}

	})
}
