package handler

import (
	"fmt"
	"github/edrpereira_meli/go-web/products"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Published   bool    `json:"published"`
	DateCreated string  `json:"dateCreated"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{service: p}
}

func (p *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}

func (p *ProductHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id := ctx.Param("id")
		p1, err := p.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf(err.Error())})
			return
		}
		ctx.JSON(http.StatusOK, p1)
	}
}

func (p *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ps, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Published, req.DateCreated)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, ps)
	}
}

func (p *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Inválido Id"})
			return
		}

		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		if req.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O cor do produto é obrigatório"})
			return
		}
		if req.Price == 0.0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O preço do produto é obrigatório"})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O estoque do produto é obrigatório"})
			return
		}
		if req.DateCreated == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A data de criação do produto é obrigatório"})
			return
		}

		ps, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Published, req.DateCreated)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, ps)
	}
}

func (p *ProductHandler) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if token := ctx.GetHeader("token"); token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "inválido Id"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do produto é obrigatório"})
			return
		}
		if req.Price == 0.0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do preço é obrigatório"})
			return
		}
		productToUpdate, err := p.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, productToUpdate)
	}
}

func (p *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Inválido Id"})
			return
		}
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}
