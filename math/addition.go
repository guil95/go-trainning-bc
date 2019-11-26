package math

type Addition struct {
	number1 float64
	number2 float64
}

func NewAddtion(number1 float64, number2 float64) *Addition {
	return &Addition{number1: number1, number2: number2}
}

func (a Addition) Calc() float64 {
	return a.number1 + a.number2
}

