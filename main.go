package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definition of the book type
type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

// The collection of books - In real life, It will be a database.
var books = []book{
	{ID: "1", Title: "In Search Of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// Get all books using GET request
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Gets a book by its id
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book Not Found")
}

// Calls the function and uses GET request
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// Loan a book - Substract from quantity 1 - If valid
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// id parameter Wasn't given
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	// Id Of book doesn't exists
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}

	// Book quantity is lower or equal than 0
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Not Available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// Returns a book - Add to quantity 1 - If Valid
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// id parameter Wasn't given
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := getBookById(id)

	// Id Of book doesn't exists
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// Creates a book using POST request - with validation
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Contains the routes(urls) for each end-point
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
