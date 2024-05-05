package entities

import (
	applicationCommon "clean-go/application/common"

	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_SALT int = 12

type User struct {
	Id       int
	Email    string
	Password string
}

func NewUser(email, password string) *User {
	p := new(User)

	p.Email = email
	p.Password = password

	return p
}

func (u *User) SetId(id int) {
	u.Id = id
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), BCRYPT_SALT)

	if err != nil {
		return applicationCommon.NewAppError(applicationCommon.InternalServerError, "")
	}

	u.Password = string(bytes)

	return nil
}

func (u *User) IsPasswordVerified(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
