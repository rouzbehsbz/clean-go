package usecases

import "clean-go/domain/entities"

type IUserService interface {
	Login(email, password string) error
	Register(email, password string) error
}

type IUserRepository interface {
	Create(user *entities.User) (int, error)
	FindByEmail(email string) (*entities.User, error)
	IsEmailExists(email string) (bool, error)
}
