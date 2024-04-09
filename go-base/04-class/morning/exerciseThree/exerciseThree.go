package exercise

import (
	"fmt"
)

func ExerciseThree() {
	fmt.Println("============ Exercicio 03 ============")

	salary := 10
	s2 := 100000

	formtMsg(verifySalary, salary)
	formtMsg(verifySalary, s2)

}

func Error(s int) error {
	return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", s)
}

func verifySalary(s int) (string, error) {
	if s < 15000 {
		return "", Error(s)
	}
	return "Necessário pagamento de imposto", nil
}

func formtMsg(verify func(a int) (string, error), s int) {
	if res, err := verify(s); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
