package books_service

import (
	"goRestServices/pkg/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	books := books.FindBooks(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": books})
}

func CreateBook(c *gin.Context) {
	book := books.CreateBook(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": book})
}

func FindBook(c *gin.Context) {
	book := books.FindBook(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": book})
}

func UpdateBook(c *gin.Context) {
	book := books.UpdateBook(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": book})
}

func DeleteBook(c *gin.Context) {
	r := books.DeleteBook(c)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": r})
}
