package exercise

import "fmt"

func ExerciseFour() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("Benjamin tem %d anos.\n", employees["Benjamin"])

	totalMoreTwentyOne := 0

	for _, employee := range employees {
		if employee > 21 {
			totalMoreTwentyOne += 1
		}
	}

	fmt.Printf("Total de funcionários com mais de 21 anos: %d\n", totalMoreTwentyOne)

	employees["Federico"] = 25
	delete(employees, "Pedro")
}
