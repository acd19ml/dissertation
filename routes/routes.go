package routes

import (
	"net/http"

	. "github.com/acd19ml/dissertation/database"
	. "github.com/acd19ml/dissertation/models"
	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	ListBooks := List_Books()
	if ListBooks == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No books found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Books found", "data": ListBooks})
}

func FindBook(c *gin.Context) {
	id := c.Param("id")
	book := Find_Book(id)
	if book == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book found", "data": book})
}

func CreateBook(c *gin.Context) {
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	id := Create_Book(book)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book created", "data": id})
}
