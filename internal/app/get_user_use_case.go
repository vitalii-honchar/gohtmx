package app

import "go-htmx/internal/domain"

type (
	GetUserUseCase interface {
		GetUser(id int) (*domain.User, error)
	}

	UserStorage interface {
		GetUser(id int) (*domain.User, error)
	}
)

type getUserUseCase struct {
	storage UserStorage
}

func NewGetUserUseCase(userStorage UserStorage) GetUserUseCase {
	return &getUserUseCase{
		storage: userStorage,
	}
}

func (u *getUserUseCase) GetUser(id int) (*domain.User, error) {
	return u.storage.GetUser(id)
}
