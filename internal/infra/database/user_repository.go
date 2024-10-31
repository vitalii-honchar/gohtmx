package database

import "go-htmx/internal/domain"

type UserRepository struct {
	storage map[int]*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		storage: map[int]*domain.User{
			1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
			2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
		},
	}
}

func (r *UserRepository) GetUser(id int) (*domain.User, error) {
	return r.storage[id], nil
}
