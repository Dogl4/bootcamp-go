package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          string  `form:"id" json:"id"`
	Name        string  `form:"name" json:"name"`
	Color       string  `form:"color" json:"color"`
	Price       float64 `form:"price" json:"price"`
	Stock       int     `form:"stock" json:"stock"`
	Published   bool    `form:"published" json:"published"`
	DateCreated string  `form:"dateCreated" json:"dateCreated"`
}

var ps []Product
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store()
}

func GetAll(ctxt *gin.Context) {
	ctxt.JSON(http.StatusOK, ps)
}

func GetProductById(ctxt *gin.Context) {
	id := ctxt.Param("id")

	for _, item := range products {
		if item.Id == id {
			ctxt.JSON(http.StatusOK, item)
			return
		}
	}
	ctxt.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}
