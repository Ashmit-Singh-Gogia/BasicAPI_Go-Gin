package main

import (
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

func main() {
	router := gin.Default()
}
