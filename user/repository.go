package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllUser() ([]User, error)
	FindUserById(ID string) (User, error)
	CreateNewUser(user User) (User, error)
	DeleteUser(ID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllUser() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindUserById(ID string) (User, error) {
	var user User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *repository) CreateNewUser(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) DeleteUser(ID string) error {
	var user User

	err := r.db.Delete(&user, ID)

	if err != nil {
		fmt.Println("User Not Found with ID: ", ID)
	}

	return nil
}
