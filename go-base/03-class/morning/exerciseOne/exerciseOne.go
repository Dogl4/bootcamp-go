package exercise

import (
	"errors"
	"fmt"
	"os"
)

func ExerciseOne() {
	fmt.Println("============ Exercicio 01 ============")

	products := []product{
		{
			3,
			5.89,
			5,
		},
		{
			4,
			8.98,
			2,
		},
		{
			5,
			10.50,
			5,
		},
	}

	writeCsv("./exerciseOne/productClear.csv", products)
}

type product struct {
	id       int
	price    float64
	quantity int
}

func writeCsv(path string, products []product) error {
	if len(products) == 0 {
		return errors.New("Erro: quantidade inválida")
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return errors.New("Erro: ao abrir o arquivo")
	}
	defer file.Close()

	p := products[0]

	if _, err = file.WriteString(p.generateHead()); err != nil {
		return fmt.Errorf("Erro: ao gerar cabeçalho: %w", err)
	}

	for _, p := range products {
		if _, err = file.WriteString(p.generateLine()); err != nil {
			return fmt.Errorf("Erro: ao salvar linha %w", err)
		}
	}

	return nil
}

func (p product) generateHead() string {
	return "ID,Quantity,Price\n"
}

func (p product) generateLine() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantity, p.price)
}
