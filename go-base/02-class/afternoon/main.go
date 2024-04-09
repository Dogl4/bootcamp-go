package main

import (
	exerciseOne "exercise/exerciseOne"
	exerciseTwo "exercise/exerciseTwo"
	"fmt"
)

func main() {
	exerciseOne.ExerciseOne(
		exerciseOne.Student{"Pedro", "Barreto", 12345678, exerciseOne.Date{2024, 4, 2}},
	)
	fmt.Println("\nExercicio 2:")
	exerciseTwo.ExerciseTwo()
}
