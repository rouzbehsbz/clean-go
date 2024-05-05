package common

import (
	usecases "clean-go/application/usecases/user"
	"clean-go/infrastructure/persistence/memory/repositories"
)

type Container struct {
	UserService    usecases.IUserService
	UserRepository usecases.IUserRepository
}

func NewContainer() *Container {
	p := new(Container)

	p.UserRepository = repositories.NewUserRepository()
	p.UserService = usecases.NewUserService(p.UserRepository)

	return p
}
