package controllers

import (
	"golang-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorController struct {
	DB *gorm.DB
}

func NewAuthorController(db *gorm.DB) *AuthorController {
	return &AuthorController{DB: db}
}

func (bc *AuthorController) CreateAuthor(c *gin.Context) {
	var requestAuthor models.CreateAuthorRequest
	var author models.Author

	if err := c.ShouldBindJSON(&requestAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author = requestAuthor.ToModel()

	// Creating new author, preload the author data before returning
	result := bc.DB.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating author"})
		return
	}

	// Preload the author data before returning
	bc.DB.Preload("Author").First(&author, author.ID)

	c.JSON(http.StatusCreated, author)
}
