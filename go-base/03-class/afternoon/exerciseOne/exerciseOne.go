package exercise

import "fmt"

func ExerciseOne() {
	fmt.Println("============ Exercicio 01 ============")

	p1 := user{
		"Pedro",
		"Barreto",
		27,
		"e@e.e",
		"123",
	}

	fmt.Printf("%v %v\n", p1.name, p1.lastName)
	p1.changeNameAndLasName("Douglas", "Pereira")
	fmt.Printf("%v %v\n", p1.name, p1.lastName)

	fmt.Printf("%v\n", p1.age)
	p1.changeAge(28)
	fmt.Printf("%v\n", p1.age)

	fmt.Printf("%v\n", p1.email)
	p1.changeEmail("a@a.c")
	fmt.Printf("%v\n", p1.email)

	fmt.Printf("%v\n", p1.password)
	p1.changePassword("456")
	fmt.Printf("%v\n", p1.password)
}

type user struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (u *user) changeNameAndLasName(name string, lasName string) {
	u.name = name
	u.lastName = lasName
}

func (u *user) changeAge(age int) {
	u.age = age
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u *user) changePassword(password string) {
	u.password = password
}
