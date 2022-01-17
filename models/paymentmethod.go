package models

type PaymentMethod interface {
	Pay(amount float32)
}
