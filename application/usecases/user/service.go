package usecases

import (
	applicationCommon "clean-go/application/common"
	"clean-go/domain/entities"
)

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	p := new(UserService)

	p.userRepository = userRepository

	return p
}

func (u *UserService) Register(email, password string) error {
	isEmailExists, err := u.userRepository.IsEmailExists(email)

	if err != nil || isEmailExists {
		return applicationCommon.NewAppError(applicationCommon.EntityExists, "This email is already exists.")
	}

	user := entities.NewUser(email, password)

	err = user.HashPassword()

	if err != nil {
		return err
	}

	userId, err := u.userRepository.Create(user)

	if err != nil {
		return err
	}

	user.SetId(userId)

	return nil
}

func (u *UserService) Login(email, password string) error {
	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return err
	}

	isPasswordVerified := user.IsPasswordVerified(password)

	if !isPasswordVerified {
		return applicationCommon.NewAppError(applicationCommon.EntityExists, "Email or password is incorrect.")
	}

	return nil
}
