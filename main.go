package main

import (
	"fmt"
	"golang-gin-gorm/controllers"
	"golang-gin-gorm/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		"5432",
		"postgres",
		"root",
		"bookstore",
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto migrate the schema
	db.AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.Category{},
	)

	// Initialize router
	router := gin.Default()

	// initialize controller
	bookController := controllers.NewBookController(db)
	authorController := controllers.NewAuthorController(db)

	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			books.GET("/", bookController.GetBooks)
			books.GET("/:id", bookController.GetBook)
			books.POST("/", bookController.CreateBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}

		author := v1.Group("/author")
		{
			author.POST("/", authorController.CreateAuthor)
		}
	}

	router.Run(":8080")
}
