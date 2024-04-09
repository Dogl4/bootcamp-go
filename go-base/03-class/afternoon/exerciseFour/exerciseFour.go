package exercise

import (
	"fmt"
	"math/rand"
	"time"
)

func ExerciseFour() {
	fmt.Println("============ Exercicio 04 ============")

	a1 := rand.Perm(100)
	a2 := rand.Perm(1000)
	a3 := rand.Perm(10000)

	fmt.Println("Com 100")
	exFunc(insertionSort, a1, "Isertion sort")
	exFunc(bubbleSort, a1, "Bubble sort")
	exFunc(selectionSort, a1, "Selection sort")

	fmt.Println("\nCom 1000")
	exFunc(insertionSort, a2, "Isertion sort")
	exFunc(bubbleSort, a2, "Bubble sort")
	exFunc(selectionSort, a2, "Selection sort")

	fmt.Println("\nCom 10000")
	exFunc(insertionSort, a3, "Isertion sort")
	exFunc(bubbleSort, a3, "Bubble sort")
	exFunc(selectionSort, a3, "Selection sort")
}

// https://blog.boot.dev/golang/insertion-sort-golang/
func insertionSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		for j := i; j > 0 && items[j-1] > items[j]; j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
	return items
}

// https://www.tutorialspoint.com/bubble-sort-in-go-lang
func bubbleSort(items []int) []int {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

// https://www.golangprograms.com/golang-program-for-implementation-of-selection-sort.html
func selectionSort(items []int) []int {
	var minIndex int
	var temp int
	for i := 0; i < len(items); i++ {
		minIndex = i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		temp = items[i]
		items[i] = items[minIndex]
		items[minIndex] = temp
	}
	return items
}

func exFunc(funcCurrent func(items []int) []int, list []int, method string) {
	startTime := time.Now()

	funcCurrent(list)

	elapsed := time.Since(startTime)
	fmt.Printf("Metodo: %v. Tempo decorrido: %v\n", method, elapsed)
}
