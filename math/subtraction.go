package math

type Subtraction struct {
	number1 float64
	number2 float64
}

func NewSubtraction(number1 float64, number2 float64) *Subtraction {
	return &Subtraction{number1: number1, number2: number2}
}

func (s Subtraction) Calc() float64 {
	return s.number1 - s.number2
}
