package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func loadProjectRoutes(group *echo.Group) {
	defaultAddressMessage(group)
}

type defaultResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func handlerFunc(c echo.Context) error {
	defaultResponse := defaultResponse{
		StatusCode: http.StatusOK,
		Message: "Seja bem-vindo(a) a nossa API REST",
	}

	return c.JSON(http.StatusOK, defaultResponse)
}

func loadDefaultMessage() echo.HandlerFunc {
	return handlerFunc
}

func defaultAddressMessage(group *echo.Group) {
	group.GET("", loadDefaultMessage())
}
