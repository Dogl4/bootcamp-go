package exercise

import (
	"errors"
	"fmt"
)

func ExerciseThree(hours float64, category string) {
	result, err := calculationSalary(hours, category)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O Salário é: R$ %.2f\n", result)
	}
}

func calculationSalary(hours float64, category string) (float64, error) {
	switch {
	case category == "C":
		return hours * 1000.00, nil
	case category == "B":
		return hours * 1500.00, nil
	case category == "A":
		return hours * 3000.00, nil
	default:
		return 0, errors.New("Categoria inválida.")
	}
}
