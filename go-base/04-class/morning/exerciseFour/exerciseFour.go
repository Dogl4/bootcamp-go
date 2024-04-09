package exercise

import (
	"errors"
	"fmt"
)

func ExerciseFour() {
	fmt.Println("============ Exercicio 04 ============")

	formtMsg(setSalary, 81, 20)
	formtMsg(setSalary, 79, 19)
	formtMsg(setSalary, 100, 1000)
}

func setSalary(hour int, valueOfHour int) (float64, error) {
	if hour < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	salary := float64(hour) * float64(valueOfHour)
	if salary >= 20000 {
		tax := salary * 0.1
		salary -= tax
	}
	return salary, nil
}

func Error(s int) error {
	return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", s)
}

func formtMsg(verify func(a int, b int) (float64, error), hour int, valueOfHour int) {
	if res, err := verify(hour, valueOfHour); err != nil {
		if errors.Unwrap(err) != nil {
			fmt.Println(errors.Unwrap(err).Error())
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Salário mensal: R$ %.2f \n", res)
	}
}
