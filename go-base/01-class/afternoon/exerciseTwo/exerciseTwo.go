package exercise

import "fmt"

func ExerciseTwo() {
	clients := []map[string]interface{}{
		{
			"name":    "Pedro",
			"age":     27,
			"salary":  110000,
			"wordAge": 3,
		},
		{
			"name":    "Lais",
			"age":     21,
			"salary":  210000,
			"wordAge": 1,
		},
		{
			"name":    "Carlos",
			"age":     25,
			"salary":  210000,
			"wordAge": 1,
		},
		{
			"name":    "Maria",
			"age":     25,
			"salary":  90000,
			"wordAge": 2,
		},
	}

	for _, client := range clients {
		fmt.Println("Cliente:", client["name"])

		if client["age"].(int) < 23 {
			fmt.Println("Empréstimo: indisponivel (menor de 23 anos)")
			continue
		}

		if client["wordAge"].(int) < 2 {
			fmt.Println("Empréstimo: indisponivel (menos de um ano de serviço)")
			continue
		}

		if client["salary"].(int) > 100000 {
			fmt.Println("Empréstimo: disponivel (sem juros)")
		} else {
			fmt.Println("Empréstimo: disponivel (com juros)")
		}
	}
}
