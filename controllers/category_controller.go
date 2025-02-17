package controllers

import (
	"golang-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (bc *CategoryController) CreateCategory(c *gin.Context) {
	var requestCategory models.CreateCategoryRequest
	var category models.Category

	if err := c.ShouldBindJSON(&requestCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category = requestCategory.ToModel()

	// Creating new category, preload the category data before returning
	result := bc.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Creating category"})
		return
	}

	// Preload the category data before returning
	bc.DB.Preload("Category").First(&category, category.ID)

	c.JSON(http.StatusCreated, category)
}
