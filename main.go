package main

import (
	"fmt"
	"log"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database")
	}

	fmt.Println("connected to the database")
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PUSTAKA API"})
	})

	v1 := router.Group("/v1")
	v1.GET("/book", bookHandler.RootHandler)
	v1.GET("/book/:id", bookHandler.BookHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/book", bookHandler.PostBookHandler)

	router.Run()
}

// main
// handler
// service
// repository
// db
