package main

import (
	exerciseFour "exercise/exerciseFour"
	exerciseOne "exercise/exerciseOne"
	exerciseThree "exercise/exerciseThree"
	exerciseTwo "exercise/exerciseTwo"
	"fmt"
)

func main() {
	exerciseOne.ExerciseOne()
	fmt.Println()
	exerciseTwo.ExerciseTwo()
	fmt.Println()
	fmt.Println(exerciseThree.ExerciseThree(7))
	fmt.Println()
	exerciseFour.ExerciseFour()
}
