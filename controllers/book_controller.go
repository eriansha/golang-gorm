package controllers

import (
	"golang-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) GetBooks(c *gin.Context) {
	var books []models.Book
	result := bc.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bc.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book
	var requestBook models.CreateBookRequest

	if err := c.ShouldBindJSON(&requestBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book = requestBook.ToModel()

	result := bc.DB.Create(&book).Preload("Author")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating book"})
		return
	}

	// Preload the author data before returning
	bc.DB.Joins("Author").First(&book, book.ID)

	c.JSON(http.StatusCreated, book)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := bc.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bc.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	result := bc.DB.Delete(&models.Book{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
