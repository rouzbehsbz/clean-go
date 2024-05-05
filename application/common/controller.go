package applicationCommon

import "github.com/labstack/echo/v4"

func NewControllerHandler(co echo.Context, response *ApiResponse) error {
	return co.JSON(response.Code, response)
}
