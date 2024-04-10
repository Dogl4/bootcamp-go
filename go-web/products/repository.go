package products

import (
	"fmt"
)

type Product struct {
	Id          int
	Name        string
	Color       string
	Price       float64
	Stock       int
	Published   bool
	DateCreated string
}

var productsDataBase []Product
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error)
	LastId() int
	Update(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return productsDataBase, nil
}

func (r *repository) LastId() int {
	lastId++
	return lastId
}

func (r *repository) GetById(id int) (Product, error) {
	for _, item := range productsDataBase {
		if item.Id == id {
			return item, nil
		}
	}
	return Product{}, fmt.Errorf("error: Not fount product")
}

func (r *repository) Store(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error) {
	p := Product{id, name, color, price, stock, published, dateCreated}
	productsDataBase = append(productsDataBase, p)
	return p, nil
}

func (r *repository) Update(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Published: published, DateCreated: dateCreated}
	updated := false
	for i := range productsDataBase {
		if productsDataBase[i].Id == id {
			p.Id = id
			productsDataBase[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d não encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	var productToUpdate Product
	updated := false
	for i := range productsDataBase {
		if productsDataBase[i].Id == id {
			productsDataBase[i].Name = name
			productsDataBase[i].Price = price
			updated = true
			productToUpdate = productsDataBase[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Produto %d não encontrado", id)
	}
	return productToUpdate, nil
}

func (r *repository) Delete(id int) error {
	delete := false
	var index int
	for i := range productsDataBase {
		if productsDataBase[i].Id == id {
			index = i
			delete = true
		}
	}
	if !delete {
		return fmt.Errorf("Produto %d não encontrado", id)
	}
	productsDataBase = append(productsDataBase[:index], productsDataBase[index+1:]...)
	return nil
}
