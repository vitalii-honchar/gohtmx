package database

import "go-htmx/internal/domain"

type userRepository struct {
	storage map[int]*domain.User
}

func newUserRepository() *userRepository {
	return &userRepository{
		storage: map[int]*domain.User{
			1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
			2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
		},
	}
}

func (r *userRepository) GetUser(id int) (*domain.User, error) {
	return r.storage[id], nil
}
