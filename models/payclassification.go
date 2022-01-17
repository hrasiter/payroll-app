package models

type PayClassification interface {
	CalculatePay() float32
}
