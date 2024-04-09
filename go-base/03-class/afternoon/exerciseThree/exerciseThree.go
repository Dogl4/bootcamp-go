package exercise

import (
	"fmt"
	"os"
	"time"
)

func ExerciseThree() {
	fmt.Println("\n============ Exercicio 03 ============")

	p1 := []product{{"teclado", 220.90, 2}}
	s1 := []service{{"Assistência", 10.0, 120}}
	m1 := []maintenance{{"Formatar pc", 200.00}}

	go sumProduct(p1)
	go sumService(s1)
	go sumMaintenance(m1)
	exitScanln()
}

type product struct {
	name     string
	price    float64
	quantity int
}

type service struct {
	name      string
	price     float64
	minOfWork float64
}

type maintenance struct {
	name  string
	price float64
}

func sumProduct(products []product) (total float64) {
	total = 0.0
	for _, p := range products {
		total += p.price * float64(p.quantity)
	}
	fmt.Printf("Total do preço de %v é R$ %.2f\n", "produtos", total)
	return
}

func sumService(services []service) (total float64) {
	total = 0.0
	for _, s := range services {
		total += s.price * s.minOfWork
	}
	fmt.Printf("Total do preço de %v é R$ %.2f\n", "Serviços", total)
	return
}

func sumMaintenance(maintenances []maintenance) (total float64) {
	total = 0.0
	for _, m := range maintenances {
		total += m.price
	}
	fmt.Printf("Total do preço de %v é R$ %.2f\n", "Manutenção", total)
	return
}

func exitScanln() {
	ch := make(chan int)

	go func() {
		fmt.Scanf("\n")
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("Exiting.")
		os.Exit(0)
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out, exiting.")
	}
}
