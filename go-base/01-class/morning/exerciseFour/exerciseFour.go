package exerciseFour

import "fmt"

func ExerciseFour() {
	var sobrenome string = "Silva"
	var idade int = 25
	var boolean bool = false
	var salario float64 = 4585.90
	var nome string = "Fellipe"

	fmt.Println("Nome: ", nome)
	fmt.Println("Sobrenome: ", sobrenome)
	fmt.Println("Idade: ", idade)
	fmt.Println("Boolean: ", boolean)
	fmt.Printf("Salário: R$ %.2f ", salario)
}
