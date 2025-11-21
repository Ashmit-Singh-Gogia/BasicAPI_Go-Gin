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

func getBookById(c *gin.Context) {
	id := c.Param("id")
	for i := range books {
		if books[i].ID == id {
			c.IndentedJSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Book with id not found",
	})
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func createBook(c *gin.Context) {
	var newBook Book
	err := c.BindJSON(&newBook)
	if err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}
func checkOut(c *gin.Context) {
	id := c.Query("id")
	for i := range books {
		if books[i].ID == id && books[i].Quantity > 0 {
			books[i].Quantity -= 1
			c.IndentedJSON(http.StatusOK, books[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"error": "Book not found or out of stock",
	})
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.GET("/books/checkout", checkOut)
	router.POST("/books", createBook)
	router.Run(":8080")
}
