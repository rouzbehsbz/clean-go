package common

import (
	usecases "clean-go/application/usecases/user"
	"clean-go/infrastructure/persistence/postgres"
	"clean-go/infrastructure/persistence/postgres/repositories"
)

type Container struct {
	UserService    usecases.IUserService
	UserRepository usecases.IUserRepository
}

func NewContainer(config *Config) (*Container, error) {
	p := new(Container)

	postgresDb, err := postgres.GetInstance(config.GetDatabaseConnectionString())

	if err != nil {
		return nil, err
	}

	p.UserRepository = repositories.NewUserRepository(postgresDb)
	p.UserService = usecases.NewUserService(p.UserRepository)

	return p, nil
}
