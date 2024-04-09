package products

type Service interface {
	GetAll() ([]Product, error)
}
