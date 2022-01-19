package models

import (
	"testing"
)

func Test_TimecardTransaction(t *testing.T) {
	t.Run("Add timecard to hourly employee", func(t *testing.T) {
		empId := 2
		tr := NewAddHourlyEmployeeTransaction(empId, "noname", 1.5)

		tr.Execute()
		tct := NewTimecardTransaction(8, "20220119", empId)
		tct.Execute()
		e := PayrollDB.GetEmployee(empId)
		pc := e.GetClassification()
		hc, ok := pc.(*HourlyClassification)

		if !ok {
			t.Errorf("Employee classification is not hourly")
		}

		tc := hc.GetTimeCard("20220119")

		if tc.GetHours() != 8 {
			t.Errorf("want: %d, got: %d", 8, tc.GetHours())
		}

	})

	t.Run("Add timecard to hourly employee", func(t *testing.T) {
		empId := 2

		tr := NewAddSalariedEmployeeTransaction(empId, "noname", 1500)
		tr.Execute()
		tct := NewTimecardTransaction(8, "20220119", empId)
		tct.Execute()
		e := PayrollDB.GetEmployee(empId)
		pc := e.GetClassification()
		hc, ok := pc.(*HourlyClassification)

		if !ok {
			t.Errorf("Employee classification is not hourly")
		}

		tc := hc.GetTimeCard("20220119")

		if tc.GetHours() != 8 {
			t.Errorf("want: %d, got: %d", 8, tc.GetHours())
		}

	})
}
