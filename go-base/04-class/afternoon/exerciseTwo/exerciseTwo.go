package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PATH_FILE = "./exerciseTwo/customers.txt"
)

func ExerciseTwo() {
	fmt.Println("============ Exercicio 02 ============")

	c1 := client{
		name:     "Pedro",
		lastName: "Douglas",
		rg:       "123456",
		phone:    "11 9999 9999",
		address:  "Rua Centro",
	}

	saveClient(c1)
}

type client struct {
	id       int
	name     string
	lastName string
	rg       string
	phone    string
	address  string
}

func readClients() ([]client, error) {
	file, err := os.Open(PATH_FILE)
	if err != nil {
		return nil, fmt.Errorf("error: o arquivo indicado não foi econtrado ou está danificado")
	}
	defer file.Close()

	var clients []client
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 6 {
			return nil, fmt.Errorf("error: o arquivo indicado não está no formato correto")
		}
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}

		client := client{
			id:       id,
			name:     fields[1],
			lastName: fields[2],
			rg:       fields[3],
			phone:    fields[4],
			address:  fields[5],
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func generateId() (int, error) {
	clients, err := readClients()
	if err != nil {
		return 0, err
	}
	return len(clients) + 1, nil
}

func saveFile(clients []client, id int) error {
	file, err := os.Create(PATH_FILE)
	if err != nil {
		return fmt.Errorf("erro: ao criar o arquivo")
	}
	defer file.Close()

	for i := 0; i < len(clients); i += 1 {
		breakLine := "%d,%s,%s,%s,%s,%s\n"
		if i == len(clients)-1 {
			breakLine = "%d,%s,%s,%s,%s,%s"
		}
		c := clients[i]

		idCurrent := c.id
		if idCurrent == c.id-1 {
			idCurrent = id
		}

		_, err := fmt.Fprintf(file, breakLine, idCurrent, c.name, c.lastName, c.rg, c.phone, c.address)
		if err != nil {
			return fmt.Errorf("erro: ao escrever o arquivo")
		}
	}
	return nil
}

func saveClient(c client) {
	id, err := generateId()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(PATH_FILE)
	if err != nil {
		panic(fmt.Errorf("error: ao ler aquivo"))
	}

	clients, err := readClients()
	if err != nil {
		panic(err)
	}

	clients = append(clients, client{
		id:       id,
		name:     c.name,
		lastName: c.lastName,
		rg:       c.rg,
		phone:    c.phone,
		address:  c.address,
	})

	err = saveFile(clients, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucesso ao adicionar um novo cliente")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado: ", r)
		}
	}()
}
