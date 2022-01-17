package models

type PaymentSchedule interface {
	IsPayDay() bool
}
