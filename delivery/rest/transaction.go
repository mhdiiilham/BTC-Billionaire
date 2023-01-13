package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mhdiiilham/BTC-Billionaire/model"
	"github.com/mhdiiilham/BTC-Billionaire/service"
	"github.com/sirupsen/logrus"
)

func StoreNewTransactionHandler(transactioner Transactioner) echo.HandlerFunc {
	return func(c echo.Context) error {
		logrus.Info("received http request to store new transaction")

		var req model.TransactionRequest
		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			return c.JSON(http.StatusBadRequest, model.HTTPResponse{
				StatusCode: http.StatusBadRequest,
				Message:    http.StatusText(http.StatusBadRequest),
				Data:       nil,
				Error:      err.Error(),
			})
		}

		ctx := context.Background()
		if err := transactioner.RecordNewTransaction(ctx, req.Datetime, req.Amount); err != nil {
			var (
				errResp    = errors.New("internal server error")
				resp       model.HTTPResponse
				statusCode int
			)

			switch {
			case errors.Is(err, service.ErrInvalidDateTimeFormat) || errors.Is(err, service.ErrInvalidAmount):
				statusCode = http.StatusBadRequest
				errResp = err
			default:
				statusCode = http.StatusInternalServerError
			}

			resp = model.HTTPResponse{
				StatusCode: statusCode,
				Message:    http.StatusText(statusCode),
				Data:       nil,
				Error:      errResp.Error(),
			}
			return c.JSON(statusCode, resp)
		}

		resp := model.HTTPResponse{StatusCode: http.StatusCreated, Message: http.StatusText(http.StatusCreated), Data: req}
		return c.JSON(http.StatusCreated, resp)
	}
}

func GetBalanceHistoryHandler(transactioner Transactioner) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request model.GetBalanceHistoryRequest

		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			return c.JSON(http.StatusBadRequest, model.HTTPResponse{
				StatusCode: http.StatusBadRequest,
				Message:    http.StatusText(http.StatusBadRequest),
				Data:       nil,
				Error:      err.Error(),
			})
		}

		ctx := context.Background()
		balanceHistories, err := transactioner.GetBalanceHistory(ctx, request.StartDateTime, request.EndDateTime)
		if err != nil {
			var (
				errResp    = errors.New("internal server error")
				resp       model.HTTPResponse
				statusCode int
			)

			switch {
			case errors.Is(err, service.ErrInvalidDateTimeFormat):
				statusCode = http.StatusBadRequest
				errResp = err
			default:
				statusCode = http.StatusInternalServerError
			}

			resp = model.HTTPResponse{
				StatusCode: statusCode,
				Message:    http.StatusText(statusCode),
				Data:       nil,
				Error:      errResp.Error(),
			}
			return c.JSON(statusCode, resp)
		}

		resp := model.HTTPResponse{StatusCode: http.StatusOK, Message: http.StatusText(http.StatusOK), Data: balanceHistories}
		return c.JSON(resp.StatusCode, resp)

	}
}
