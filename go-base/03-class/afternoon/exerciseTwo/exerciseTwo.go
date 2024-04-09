package exercise

import "fmt"

func ExerciseTwo() {
	fmt.Println("\n============ Exercicio 02 ============")

	p1 := newProduct("teclado", 210.50)
	u1 := user{}
	u1.addProduct(p1, 2)

	fmt.Printf("Usuário: %++v\n", u1)
	u1.deleteProduct()
	fmt.Printf("Usuário: %++v\n", u1)
}

type user struct {
	name     string
	lastName string
	email    string
	products []product
}

type product struct {
	name     string
	price    float64
	quantity int
}

func newProduct(name string, price float64) product {
	return product{
		name,
		price,
		0,
	}
}

func (u *user) addProduct(p product, quantity int) {
	u.products = append(u.products, product{
		name:     p.name,
		price:    p.price,
		quantity: quantity,
	})
}

func (u *user) deleteProduct() {
	u.products = []product{}
}
