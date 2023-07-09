package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     int32  `gorm:"primaryKey,autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int32  `json:"year"`
}

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Book Library API"})
}

func NewBookHandler(c *gin.Context) {
	var newProduct Book
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db, err := DbConn()
	if err != nil {
		panic(err.Error())
	}
	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func GetBookHandler(c *gin.Context) {
	db, err := DbConn()
	if err != nil {
		panic(err.Error())
	}
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func InitHandler(c *gin.Context) {
	var products = []Book{
		{
			ID:     1,
			Title:  "Book A",
			Author: "Author A",
			Year:   1998,
		},
		{
			ID:     2,
			Title:  "Book B",
			Author: "Author B",
			Year:   2022,
		},
		{
			ID:     3,
			Title:  "Book C",
			Author: "Author C",
			Year:   1999,
		},
	}

	db, err := DbConn()
	db.AutoMigrate(&Book{})
	if err != nil {
		panic(err.Error())
	}

	db.Create(&products)
	c.JSON(http.StatusOK, gin.H{
		"message": "Init book",
	})
}

func DbConn() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	return db, nil
}

func main() {

	router := gin.Default()
	router.GET("/", HomepageHandler)
	router.GET("/init", InitHandler)
	router.GET("/books", GetBookHandler)
	router.POST("/books", NewBookHandler)
	err := router.Run(":8888")
	if err != nil {
		panic("Failed run server")
	}
}
