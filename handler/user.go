package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	users, err := h.userService.FindAllUser()

	if err != nil {
		fmt.Println("error while get all users", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "something wrong",
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success get all users",
		"data":    users,
	})
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.FindUserById(id)

	if err != nil {
		fmt.Println("error while get user id : %s", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "something wrong",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success get user by id",
		"data":    user,
	})
}

func (h *userHandler) CreateNewUser(c *gin.Context) {
	var user user.UserRequest

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		errMsg := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorsMessage := fmt.Sprintf("Error on field %s, condition : %s, ", e.Field(), e.ActualTag())
			errMsg = append(errMsg, errorsMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errMsg,
		})

		return
	}

	newUser, err := h.userService.CreateNewUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "errorr",
			"message": "error while create new user",
			"data":    nil,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success created a new user",
		"data":    newUser,
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {}
