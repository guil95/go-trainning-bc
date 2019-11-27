package main

import (
	"fmt"

	"github.com/guil95/go-trainning-bc/math"
)

func main() {
	a := calculate(1,2,"-")
	b := calculateByMap(1,2,"+")

	fmt.Println(a)
	fmt.Println(b)
}

func calculate(number1 float64, number2 float64, method string) float64 {
	switch method {
	case "+":
		return math.Exec(math.NewAddtion(number1,number2))
	case "-":
		return math.Exec(math.NewSubtraction(number1,number2))
	default:
		return 0
	}
}

func calculateByMap(number1 float64, number2 float64, method string) float64 {
	operations := map[string]math.Calculator{
		"+": math.NewAddtion(number1, number2),
		"-": math.NewSubtraction(number1, number2),
	}

	return math.Exec(operations[method])
}
