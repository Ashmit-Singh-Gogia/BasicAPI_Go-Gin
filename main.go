package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"` // ID and rest things are uppercase to make them exported
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 3},
	{ID: "2", Title: "1984", Author: "George Orwell", Quantity: 5},
	{ID: "3", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run(":8080")
}
