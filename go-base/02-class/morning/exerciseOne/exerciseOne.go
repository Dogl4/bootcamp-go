package exercise

import "fmt"

func ExerciseOne() {
	fmt.Printf("Taxa para salário de R$ 70.000 é: %d%%\n", calculationTax(70000))
	fmt.Printf("Taxa para salário de R$ 30.000 é: %d%%\n", calculationTax(30000))
	fmt.Printf("Taxa para salário de R$ 170.000 é: %d%%\n", calculationTax(170000))
}

func calculationTax(salary int) int {
	switch {
	case salary >= 50000 && salary <= 150000:
		return 17
	case salary > 150000:
		return 27
	default:
		return 0
	}
}
