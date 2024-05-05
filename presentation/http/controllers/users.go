package controllers

import (
	applicationCommon "clean-go/application/common"
	"clean-go/common"

	"github.com/labstack/echo/v4"
)

type UserRegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegisterController(c echo.Context, container *common.Container) error {
	var userRegisterDto UserRegisterDto

	if err := c.Bind(&userRegisterDto); err != nil {
		return err
	}

	if err := container.UserService.Register(userRegisterDto.Email, userRegisterDto.Password); err != nil {
		return err
	}

	res := applicationCommon.ApiResponseSuccess("Your account has been created successfully.", nil)

	return applicationCommon.NewControllerHandler(c, res)
}

func UserLoginController(c echo.Context, container *common.Container) error {
	var userLoginDto UserLoginDto

	if err := c.Bind(&userLoginDto); err != nil {
		return err
	}

	if err := container.UserService.Login(userLoginDto.Email, userLoginDto.Password); err != nil {
		return err
	}

	res := applicationCommon.ApiResponseSuccess("You have successfullly logged in.", nil)

	return applicationCommon.NewControllerHandler(c, res)
}
