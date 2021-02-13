package books

import (
	"goRestServices/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) []models.Book {
	var books []models.Book
	models.DB.Find(&books)

	return books
}

func CreateBook(c *gin.Context) models.Book {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return models.Book{}
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	return book
}

func FindBook(c *gin.Context) models.Book {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Book{}
	}

	return book
}

func UpdateBook(c *gin.Context) models.Book {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return models.Book{}
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return models.Book{}
	}

	models.DB.Model(&book).Updates(input)

	return book
}

func DeleteBook(c *gin.Context) bool {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return false
	}

	models.DB.Delete(&book)

	return true
}
