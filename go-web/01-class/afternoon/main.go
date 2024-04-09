package main

import (
	"fmt"
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

type User struct {
	Id          string  `form:"id" json:"id"`
	Name        string  `form:"name" json:"name"`
	LastName    string  `form:"lastName" json:"lastName"`
	Email       string  `form:"email" json:"email"`
	Age         int     `form:"age" json:"age"`
	Height      float64 `form:"height" json:"height"`
	Active      bool    `form:"active" json:"active"`
	DateCreated string  `form:"dateCreated" json:"dateCreated"`
}

type Transaction struct {
	Id                string  `form:"id" json:"id"`
	CodeOfTransaction string  `form:"codeOfTransaction" json:"codeOfTransaction"`
	Coin              string  `form:"coin" json:"coin"`
	Value             float64 `form:"value" json:"value"`
	Sender            string  `form:"sender" json:"sender"`
	Receiver          string  `form:"receiver" json:"receiver"`
	DateCreated       string  `form:"dateCreated" json:"dateCreated"`
}

var products = []Product{
	{
		Id:          "1",
		Name:        "Panela",
		Color:       "Prata",
		Price:       70.0,
		Stock:       5,
		Published:   true,
		DateCreated: "08/04/2024",
	},
}

var users = []User{
	{
		Id:          "1",
		Name:        "Pedro",
		LastName:    "Barreto",
		Email:       "e@e.e",
		Age:         27,
		Height:      1.72,
		Active:      true,
		DateCreated: "08/04/2024",
	},
}

var transactions = []Transaction{
	{
		Id:                "1",
		CodeOfTransaction: "hfdsajkhkj34324",
		Coin:              "Real",
		Value:             70.0,
		Sender:            "Pedro",
		Receiver:          "Pedro",
		DateCreated:       "08/04/2024",
	},
}

func main() {
	server := gin.Default()
	server.GET("/", RootPage)
	server.GET("/products", GetAllProducts)
	server.GET("/products/:id", GetProductById)
	server.GET("/users/:id", GetUserById)
	server.GET("/transactions/:id", GetTransactionById)
	server.GET("/users", GetAllUsers)
	server.GET("/transactions", GetAllTransactions)
	server.GET("/hello/:name", SendMessageWithName)
	server.Run()
}

func RootPage(ctxt *gin.Context) {
	ctxt.String(http.StatusOK, "Hello Word!")
}

func GetAllProducts(ctxt *gin.Context) {
	ctxt.JSON(http.StatusOK, products)
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

func GetUserById(ctxt *gin.Context) {
	id := ctxt.Param("id")

	for _, item := range users {
		if item.Id == id {
			ctxt.JSON(http.StatusOK, item)
			return
		}
	}
	ctxt.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func GetTransactionById(ctxt *gin.Context) {
	id := ctxt.Param("id")

	for _, item := range transactions {
		if item.Id == id {
			ctxt.JSON(http.StatusOK, item)
			return
		}
	}
	ctxt.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func GetAllUsers(ctxt *gin.Context) {
	ctxt.JSON(http.StatusOK, users)
}

func GetAllTransactions(ctxt *gin.Context) {
	ctxt.JSON(http.StatusOK, transactions)
}

func SendMessageWithName(ctxt *gin.Context) {
	message := fmt.Sprint("Olá ", ctxt.Param("name"))
	ctxt.JSON(http.StatusOK, gin.H{"message": message})
}
