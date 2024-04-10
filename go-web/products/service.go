package products

import (
	"fmt"
	"strconv"
)

type Service interface {
	GetAll() ([]Product, error)
	GetById(id string) (Product, error)
	Store(name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error)
	Update(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetById(id string) (Product, error) {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Product{}, fmt.Errorf(err.Error())
	}

	p, err := s.repository.GetById(int(num))

	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s *service) Store(name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error) {
	id := s.repository.LastId()

	p, err := s.repository.Store(id, name, color, price, stock, published, dateCreated)
	if err != nil {
		return Product{}, nil
	}
	return p, nil
}

func (s *service) Update(id int, name string, color string, price float64, stock int, published bool, dateCreated string) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, published, dateCreated)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
