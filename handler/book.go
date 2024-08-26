package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookhandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookhandler {
	return &bookhandler{bookService}
}

func (h *bookhandler) RootHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		fmt.Println("Error while get all books", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "something wrong",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success get all books",
		"data":    books,
	})
}

func (h *bookhandler) BookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := h.bookService.FindByID(id)
	if err != nil {
		fmt.Printf("Error while get book by id: %s", id, err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "something wrong",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success get book by id",
		"data":    book,
	})
}

func (h *bookhandler) QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"message": "success get book by id",
		"id":      id,
		"title":   title,
	})
}

func (h *bookhandler) PostBookHandler(c *gin.Context) {
	var book book.BookRequest

	err := c.ShouldBindBodyWithJSON(&book)
	if err != nil {

		errMsg := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			erroMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errMsg = append(errMsg, erroMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMsg,
		})
		return
	}

	newBook, err := h.bookService.Create(book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "error failed to create a new book",
			"data":    nil,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success create a 	new book",
		"data":    newBook,
	})

}
