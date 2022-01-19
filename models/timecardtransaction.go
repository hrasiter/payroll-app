package models

import "fmt"

type TimeCardTransaction struct {
	hours int
	date  string
	empid int
}

func NewTimecardTransaction(hours int, date string, eid int) Transaction {
	return &TimeCardTransaction{
		hours: hours,
		date:  date,
		empid: eid,
	}
}

func (tct *TimeCardTransaction) Execute() {
	tc := TimeCard{
		hours: tct.hours,
		date:  tct.date,
	}

	e := PayrollDB.GetEmployee(tct.empid)

	pc := e.GetClassification()

	hc, ok := pc.(*HourlyClassification)

	if ok {
		hc.AddtimeCard(tc)
	} else {
		fmt.Println("Employee is not Hourly Employee")
	}
}
