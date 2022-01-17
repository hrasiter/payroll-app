package models

type SalariedClassification struct {
	salary float32
}

func (sc *SalariedClassification) CalculatePay() float32 {
	return 0
}

func (sc *SalariedClassification) GetSalary() float32 {
	return sc.salary
}

func (sc *SalariedClassification) SetSalary(salary float32) {
	sc.salary = salary
}
