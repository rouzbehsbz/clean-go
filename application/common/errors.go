package applicationCommon

import (
	"github.com/labstack/echo/v4"
)

type AppError int

const (
	InternalServerError AppError = iota
	EntityNotFound
	EntityExists
	AuthorizationFailed
)

func NewAppError(kind AppError, message string) error {
	msg := message
	code := kind.Code()

	if msg == "" {
		msg = kind.DefaultErrorMessage()
	}

	return echo.NewHTTPError(code, msg)
}

func (a AppError) DefaultErrorMessage() string {
	switch a {
	case InternalServerError:
		return "Oops ! Something went wrong"
	case EntityNotFound:
		return "Entity not found"
	case EntityExists:
		return "Entity is already exists"
	case AuthorizationFailed:
		return "Your access is restricted"
	}

	return "Oops ! Something went wrong"
}

func (a AppError) Code() int {
	switch a {
	case InternalServerError:
		return 500
	case EntityNotFound:
		return 404
	case EntityExists:
		return 409
	case AuthorizationFailed:
		return 401
	}

	return 500
}
