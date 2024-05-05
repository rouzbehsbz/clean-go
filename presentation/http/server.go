package http

import (
	applicationCommon "clean-go/application/common"
	"clean-go/common"
	"clean-go/presentation/http/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	App       *echo.Echo
	Container *common.Container
}

func NewServer(c *common.Container) *Server {
	p := new(Server)
	p.App = echo.New()
	p.Container = c

	p.App.HTTPErrorHandler = p.ErrorHandler

	p.SetRoutes()

	return p
}

func (s *Server) Listen(host string, port uint16) {
	s.App.Start(fmt.Sprintf("%s:%d", host, port))
}

func (s *Server) SetRoutes() {
	s.App.POST("/api/v1/users", func(c echo.Context) error {
		return controllers.UserRegisterController(c, s.Container)
	})
	s.App.POST("/api/v1/users/session", func(c echo.Context) error {
		return controllers.UserLoginController(c, s.Container)
	})
}

func (s *Server) ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	if !c.Response().Committed {
		c.JSON(code, applicationCommon.ApiResponseError(code, message, nil))
	}
}
