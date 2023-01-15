package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhdiiilham/BTC-Billionaire/common"
	"github.com/mhdiiilham/BTC-Billionaire/model"
)

func RouteNotFoundHandler(cfg *common.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := map[string]any{
			"name":    "BTC-Billionaire RESTful API",
			"version": cfg.Version,
		}

		resp := model.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    http.StatusText(http.StatusNotFound),
			Data:       data,
			Error:      "routes not found!",
		}

		return c.JSON(resp.StatusCode, resp)
	}
}
