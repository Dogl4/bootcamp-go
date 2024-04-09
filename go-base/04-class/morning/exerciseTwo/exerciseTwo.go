package exercise

import (
	"errors"
	"fmt"
)

func ExerciseTwo() {
	fmt.Println("============ Exercicio 02 ============")

	salary := 10
	s2 := 100000

	formtMsg(verifySalary, salary)
	formtMsg(verifySalary, s2)

}

func Error() error {
	return errors.New("erro: O salário digitao não está dentro do valor mínimo")
}

func verifySalary(s int) (string, error) {
	if s < 15000 {
		return "", Error()
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
