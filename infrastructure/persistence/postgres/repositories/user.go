package repositories

import (
	"clean-go/domain/entities"
	"clean-go/infrastructure/persistence/postgres"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository struct {
	source *postgres.PostgresDB
}

func NewUserRepository(postgresDb *postgres.PostgresDB) *UserRepository {
	p := new(UserRepository)

	p.source = postgresDb

	return p
}

func (u *UserRepository) Create(user *entities.User) (int, error) {
	dbUser := &User{
		Email:    user.Email,
		Password: user.Password,
	}

	res := u.source.Db.Create(dbUser)

	if res.Error != nil {
		return 0, res.Error
	}

	return int(dbUser.Id), nil
}

func (u *UserRepository) FindByEmail(email string) (*entities.User, error) {
	dbUser := &User{}

	res := u.source.Db.Find(dbUser, "email = ?", email)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		return nil, nil
	}

	user := &entities.User{
		Id:       int(dbUser.Id),
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}

	return user, nil
}

func (u *UserRepository) IsEmailExists(email string) (bool, error) {
	dbUser := &User{}

	res := u.source.Db.Find(dbUser, "email = ?", email)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, res.Error
		}

		return false, nil
	}

	return true, nil
}
