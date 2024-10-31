package app

import "go-htmx/internal/domain"

type UserService interface {
	GetUser(id int) (*domain.User, error)
}

type userService struct {
	storage map[int]*domain.User
}

func NewUserService() UserService {
	return &userService{
		storage: map[int]*domain.User{
			1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
			2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
		},
	}
}

func (u *userService) GetUser(id int) (*domain.User, error) {
	return u.storage[id], nil
}
