package exercise

import (
	"errors"
	"fmt"
)

func ExerciseTwo(notes ...float64) {
	result, err := getGradeAverage(notes...)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Média de notas é: %.2f\n", result)
	}
}

func getGradeAverage(notes ...float64) (float64, error) {
	var sum, average float64
	for _, note := range notes {
		if note < 0 {
			return 0, errors.New("As notas devem ser positivas")
		}
		sum += note
	}
	average = sum / float64(len(notes))
	return average, nil
}
