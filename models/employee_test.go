package models

import (
	"testing"
)

func Test_AddSalariedEmployeeTransaction(t *testing.T) {
	id := 1
	tr := NewAddSalariedEmployeeTransaction(id, "Bob", 1000.0)
	tr.Execute()
	e := PayrollDB.GetEmployee(id)

	if e.GetName() != "Bob" {
		t.Errorf("Want %q, got % q", "Bob", e.GetName())
	}

	//TODOS: reflection
	pc := e.GetClassification()
	sc, classOk := pc.(*SalariedClassification)

	if !classOk {
		t.Errorf("type of payment classification is wrong!")
	}

	if sc.GetSalary() != 1000.0 {
		t.Errorf("want %v, got %v", 1000.0, sc.GetSalary())
	}

	ps := e.GetSchedule()
	_, scheduleOk := ps.(*MonthlySchedule)

	if !scheduleOk {
		t.Error("type of schedule is different than monthly schedule")
	}

	if !scheduleOk {
		t.Error("type of schedule is different than monthly schedule")
	}

	pm := e.GetMethod()
	_, methodOk := pm.(*HoldMethod)

	if !methodOk {
		t.Error("type of payment method is different than hold method")
	}

}
