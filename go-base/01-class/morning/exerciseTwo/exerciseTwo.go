package exerciseTwo

import "fmt"

func ExerciseTwo() {
	temperatura, umidade, pressao := 27, 48, 1015.9
	fmt.Println("Cidade: Taubaté")
	fmt.Printf("Temperatura: %d °C\n", temperatura)
	fmt.Printf("Umidade: %d%%\n", umidade)
	fmt.Printf("Pressão: %.1f mb\n", pressao)
	fmt.Println()
}
