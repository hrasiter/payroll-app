package models

type AddEmployee interface {
	GetClassification() PayClassification
	GetSchedule() PaymentSchedule
}
