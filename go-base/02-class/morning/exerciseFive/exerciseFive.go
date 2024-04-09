package exercise

import (
	"errors"
	"fmt"
)

const (
	DOG       = "Cães"
	CAT       = "Gatos"
	HAMSTER   = "Hamster"
	TARANTULA = "Tarantula"
)

type animal struct {
	name                  string
	amountOfFoodPerAnimal float64
}

type animals []animal

func ExerciseFive(animalRace string, quantity int) {
	listAnimals := animals{
		{
			DOG,
			10.0,
		},
		{
			CAT,
			5.0,
		},
		{
			HAMSTER,
			0.25,
		},
		{
			TARANTULA,
			0.15,
		},
	}
	res, err := listAnimals.verifyAnimal(animalRace, quantity)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tipo %v: %v kg(s). Para %v animais(l)\n", animalRace, res, quantity)
	}
}

func calculationFood(quantityOfAnimal int, quantityOfFood float64) float64 {
	return float64(quantityOfAnimal) * quantityOfFood
}

func (a animals) verifyAnimal(race string, quantity int) (float64, error) {
	for _, animal := range a {
		if animal.name == race {
			return calculationFood(quantity, animal.amountOfFoodPerAnimal), nil
		}
	}
	return 0.0, errors.New("Erro: animal não encontado")
}
