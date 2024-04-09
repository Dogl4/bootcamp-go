package exercise

import "fmt"

func ExerciseOne() {
	word := "palavra"

	fmt.Println("Tamanho: ", len(word))

	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
