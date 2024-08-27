package user

import "github.com/google/uuid"

type Service interface {
	FindAllUser() ([]User, error)
	FindUserById(ID string) (User, error)
	CreateNewUser(user UserRequest) (User, error)
	DeleteUser(ID string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllUser() ([]User, error) {
	users, err := s.repository.FindAllUser()

	return users, err
}

func (s *service) FindUserById(ID string) (User, error) {
	user, err := s.repository.FindUserById(ID)

	return user, err
}

func (s *service) DeleteUser(ID string) error {
	err := s.repository.DeleteUser(ID)

	return err
}

func (s *service) CreateNewUser(userRequest UserRequest) (User, error) {

	ID := uuid.New().String()

	user := User{
		ID:       ID,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
	}

	newUser, err := s.repository.CreateNewUser(user)

	return newUser, err

}
