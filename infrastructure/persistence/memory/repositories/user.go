package repositories

import (
	applicationCommon "clean-go/application/common"
	"clean-go/domain/entities"
	"clean-go/infrastructure/persistence/memory"
)

type UserRepository struct {
	nextUserId           int
	SourceIndexedById    *memory.Memory[int, *entities.User]
	SourceIndexedByEmail *memory.Memory[string, *entities.User]
}

func NewUserRepository() *UserRepository {
	p := new(UserRepository)

	p.nextUserId = 1
	p.SourceIndexedById = memory.NewMemory[int, *entities.User]()
	p.SourceIndexedByEmail = memory.NewMemory[string, *entities.User]()

	return p
}

func (u *UserRepository) Create(user *entities.User) (int, error) {
	userId := u.nextUserId

	u.SourceIndexedById.Add(userId, user)
	u.SourceIndexedByEmail.Add(user.Email, user)

	u.nextUserId += 1

	return userId, nil
}

func (u *UserRepository) FindByEmail(email string) (*entities.User, error) {
	user, ok := u.SourceIndexedByEmail.Get(email)

	if ok {
		return user, nil
	}

	return nil, applicationCommon.NewAppError(applicationCommon.EntityExists, "Email does not exists.")
}

func (u *UserRepository) IsEmailExists(email string) (bool, error) {
	_, ok := u.SourceIndexedByEmail.Get(email)

	if ok {
		return true, nil
	}

	return false, nil
}
