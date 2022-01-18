package models

import (
	"testing"
)

func Test_AddSalariedEmployeeTransaction(t *testing.T) {
	id := 1
	name := "Bob"
	salary := float32(1000.0)
	tr := NewAddSalariedEmployeeTransaction(id, name, salary)
	tr.Execute()
	e := PayrollDB.GetEmployee(id)

	if e.GetName() != name {
		t.Errorf("Want %q, got % q", name, e.GetName())
	}

	//TODOS: reflection
	pc := e.GetClassification()
	sc, classOk := pc.(*SalariedClassification)

	if !classOk {
		t.Errorf("type of payment classification is wrong!")
	}

	if sc.GetSalary() != salary {
		t.Errorf("want %v, got %v", salary, sc.GetSalary())
	}

	ps := e.GetSchedule()
	_, scheduleOk := ps.(*MonthlySchedule)

	if !scheduleOk {
		t.Error("type of schedule is different than monthly schedule")
	}

	pm := e.GetMethod()
	_, methodOk := pm.(*HoldMethod)

	if !methodOk {
		t.Error("type of payment method is different than hold method")
	}

}

func Test_AddHourlyEmployeeTransaction(t *testing.T) {
	t.Run("Normal Scenario", func(t *testing.T) {
		id := 1
		name := "skirtel"
		hourlyrate := float32(1.1)
		tr := NewAddHourlyEmployeeTransaction(id, name, hourlyrate)
		tr.Execute()
		e := PayrollDB.GetEmployee(id)

		if e.GetName() != name {
			t.Errorf("Want %q, got % q", name, e.GetName())
		}

		//TODOS: reflection
		pc := e.GetClassification()
		sc, classOk := pc.(*HourlyClassification)

		if !classOk {
			t.Errorf("type of payment classification is wrong!")
		}

		if sc.GetHourlyRate() != hourlyrate {
			t.Errorf("want %v, got %v", hourlyrate, sc.GetHourlyRate())
		}

		ps := e.GetSchedule()
		_, scheduleOk := ps.(*WeeklySchedule)

		if !scheduleOk {
			t.Error("type of schedule is different than Weekly schedule")
		}

		pm := e.GetMethod()
		_, methodOk := pm.(*HoldMethod)

		if !methodOk {
			t.Error("type of payment method is different than hold method")
		}
	})
}
