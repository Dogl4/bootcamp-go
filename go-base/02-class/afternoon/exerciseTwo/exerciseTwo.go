package exercise

import (
	"errors"
	"fmt"
)

func ExerciseTwo() {
	p1 := product{
		"Coberto",
		2,
		45.0,
	}
	p2 := product{
		"Panela",
		1,
		20.0,
	}
	newStore := &store{}

	newStore.add(p1)
	newStore.add(p2)

	res, err := newStore.total()

	if err != nil {
		fmt.Println("Batata", err)
	} else {
		fmt.Printf("Total %.2f\n", res)
	}

}

type product struct {
	name  string
	size  int
	price float64
}

type store struct {
	general []product
}
type Product interface {
	calculatePrice()
}

type Ecommerce interface {
	total() (float64, error)
	add(p product)
}

func (p product) calculatePrice(size int) (float64, error) {
	switch {
	case size == 1:
		return p.price, nil
	case size == 2:
		return p.price * 1.03, nil
	case size == 3:
		return (p.price * 1.06) + 2500.0, nil
	default:
		return 0.0, errors.New("Erro: tamanho inválido.")
	}
}
func (s store) total() (float64, error) {
	total := 0.0
	for _, product := range s.general {
		res, err := product.calculatePrice(product.size)
		if err != nil {
			return 0.0, errors.New("Erro: total loja.")
		} else {
			total += float64(res)
		}
	}
	return total, nil
}
func (s *store) add(product product) {
	s.general = append(s.general, product)
}

func newProduct(name string, size int, price float64) product {
	return product{
		name,
		size,
		price,
	}
}

func newStore() Ecommerce {
	// return append(store{}, products...)
	return &store{}
}
