package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/ishtiaqhimel/news-portal/cms/internal/db"
	systemDelivery "github.com/ishtiaqhimel/news-portal/cms/internal/system/delivery"
	systemRepository "github.com/ishtiaqhimel/news-portal/cms/internal/system/repository"
	systemUsecase "github.com/ishtiaqhimel/news-portal/cms/internal/system/usecase"
)

func Serve(stopCh <-chan struct{}) error {
	// connect to postgres DB
	if err := db.Connect(); err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	pgClient := db.Get().DB

	// http server setup
	e := echo.New()

	// repository
	sysRepo := systemRepository.NewSystemRepository(pgClient)

	// usecase
	sysUsecase := systemUsecase.NewSystemUsecase(sysRepo)

	// delivery
	systemDelivery.NewSystemHandler(e, sysUsecase)

	// start http server
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Get().App.Port)))
	}()

	// graceful shutdown setup
	<-stopCh
	logrus.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	logrus.Infof("server shutdowns gracefully")

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close db connection: %s", err)
	}

	return nil
}
