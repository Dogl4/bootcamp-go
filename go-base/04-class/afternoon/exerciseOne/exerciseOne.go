package exercise

import (
	"fmt"
	"os"
)

func ExerciseOne() {
	fmt.Println("============ Exercicio 01 ============")
	path := "./exerciseTwo/customers.txt"
	openFile(path)
}

func openFile(path string) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(fmt.Errorf("o arquivo indicado não foi encontrado ou está danificado"))
	}
	fmt.Println(*&file)
}
