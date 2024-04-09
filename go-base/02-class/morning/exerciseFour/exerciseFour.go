package exercise

import (
	"errors"
	"fmt"
)

const (
	MINIMUN = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

func ExerciseFour(typeOfOperation string, values ...float64) {
	res, err := operation(typeOfOperation, values...)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Valor é: R$ %.2f. Operação: %v\n", res, typeOfOperation)
	}
}

func operation(typeOfOperation string, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Erro: valor de entrada inválido")
	}

	switch {
	case typeOfOperation == MINIMUN:
		return opMin(values...)
	case typeOfOperation == MAXIMUM:
		return opMax(values...)
	case typeOfOperation == AVERAGE:
		return opAve(values...)
	default:
		return 0, errors.New("Erro: operação inválida")
	}
}

func opMin(values ...float64) (float64, error) {
	minValue := values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}
	return minValue, nil
}

func opMax(values ...float64) (float64, error) {
	maxValue := values[0]
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue, nil
}

func opAve(values ...float64) (float64, error) {
	sumOfValues := 0.0
	for _, value := range values {
		sumOfValues += value
	}
	return sumOfValues / float64(len(values)), nil
}
