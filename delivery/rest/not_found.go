package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhdiiilham/BTC-Billionaire/model"
)

func RouteNotFoundHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := model.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    http.StatusText(http.StatusNotFound),
			Data:       nil,
			Error:      "routes not found!",
		}

		return c.JSON(resp.StatusCode, resp)
	}
}
