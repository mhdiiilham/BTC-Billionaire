package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mhdiiilham/BTC-Billionaire/common"
	"github.com/mhdiiilham/BTC-Billionaire/delivery/rest"
	"github.com/mhdiiilham/BTC-Billionaire/repository"
	"github.com/mhdiiilham/BTC-Billionaire/service"
	"github.com/sirupsen/logrus"
)

var version = "v0.0.0"

func main() {
	cfg := common.ReadConfig(version)
	db := common.ConnectToSQL(cfg)

	// define repository
	transactionRepository := repository.NewTransactionRepository(db)

	// define service
	transactioner := service.NewTransactionService(transactionRepository)

	// Handle 404 error
	echo.NotFoundHandler = rest.RouteNotFoundHandler(cfg)

	// routing
	e := echo.New()
	apiV1 := e.Group("api/v1")

	// /api/v1/tansactions handler
	transactionsRoutes := apiV1.Group("/transactions")
	transactionsRoutes.POST("", rest.StoreNewTransactionHandler(transactioner))
	transactionsRoutes.POST("/balances", rest.GetBalanceHistoryHandler(transactioner))

	go func() {
		if err := e.Start(cfg.GetServerPort()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	logrus.Info("waiting shutdown signal")
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, done := context.WithTimeout(context.Background(), 30*time.Second)
	defer done()

	logrus.Infof("shutting down echo server; error=%v", e.Shutdown(ctx))
	logrus.Info("server is shuting down...")
	logrus.Infof("closing db connection; %v", db.Close())

}
