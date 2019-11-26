package math

type Subtration struct {
	number1 float64
	number2 float64
}

func NewSubtration(number1 float64, number2 float64) *Subtration {
	return &Subtration{number1: number1, number2: number2}
}

func (s Subtration) Calc() float64 {
	return s.number1 - s.number2
}

