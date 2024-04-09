package main

import (
	exerciseFive "exercise/exerciseFive"
	exerciseFour "exercise/exerciseFour"
	exerciseOne "exercise/exerciseOne"
	exerciseThree "exercise/exerciseThree"
	exerciseTwo "exercise/exerciseTwo"
	"fmt"
)

func main() {
	exerciseOne.ExerciseOne()
	fmt.Println()
	exerciseTwo.ExerciseTwo(3.0, 7.0, 10.0)
	exerciseTwo.ExerciseTwo(-9)
	fmt.Println()
	exerciseThree.ExerciseThree(220.0, "C")
	exerciseThree.ExerciseThree(220.0, "E")
	fmt.Println()
	exerciseFour.ExerciseFour(exerciseFour.MINIMUN, 1.0, 6.0, 7.0)
	exerciseFour.ExerciseFour(exerciseFour.MAXIMUM, 1.0, 6.0, 7.0)
	exerciseFour.ExerciseFour(exerciseFour.AVERAGE, 1.0, 6.0, 7.0)
	exerciseFour.ExerciseFour(exerciseFour.AVERAGE)
	exerciseFour.ExerciseFour("Batata", 1.0, 6.0, 7.0)
	fmt.Println()
	exerciseFive.ExerciseFive(exerciseFive.HAMSTER, 2)
	exerciseFive.ExerciseFive(exerciseFive.CAT, 2)
	exerciseFive.ExerciseFive(exerciseFive.DOG, 2)
	exerciseFive.ExerciseFive(exerciseFive.TARANTULA, 2)
	exerciseFive.ExerciseFive("Batata", 2)
}
