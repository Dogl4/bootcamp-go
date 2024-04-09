package exercise

import "fmt"

func ExerciseOne() {
	fmt.Println("============ Exercicio 01 ============")

	salary := 10
	s2 := 100000

	verifySalary(salary)
	verifySalary(s2)

}

func Error() string {
	return "erro: O salário digitao não está dentro do valor mínimo"
}

func verifySalary(s int) {
	if s < 15000 {
		fmt.Println(Error())
	}
	fmt.Println("Necessário pagamento de imposto")
}
