package server

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	articleDelivery "github.com/ishtiaqhimel/news-portal/cms/internal/article/delivery"
	articleRepository "github.com/ishtiaqhimel/news-portal/cms/internal/article/repository"
	articleUsecase "github.com/ishtiaqhimel/news-portal/cms/internal/article/usecase"
	categoryDelivery "github.com/ishtiaqhimel/news-portal/cms/internal/category/delivery"
	categoryRepository "github.com/ishtiaqhimel/news-portal/cms/internal/category/repository"
	categoryUsecase "github.com/ishtiaqhimel/news-portal/cms/internal/category/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
	"github.com/ishtiaqhimel/news-portal/cms/internal/db"
	"github.com/ishtiaqhimel/news-portal/cms/internal/middlewares"
	systemDelivery "github.com/ishtiaqhimel/news-portal/cms/internal/system/delivery"
	systemRepository "github.com/ishtiaqhimel/news-portal/cms/internal/system/repository"
	systemUsecase "github.com/ishtiaqhimel/news-portal/cms/internal/system/usecase"
	"github.com/ishtiaqhimel/news-portal/cms/internal/utils"
)

func Serve(stopCh <-chan struct{}) error {
	// connect to postgres DB
	if err := db.Connect(); err != nil {
		return fmt.Errorf("failed to connect to db: %v", err)
	}
	pgClient := db.Get().DB

	// http server setup
	e := echo.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	middlewares.Attach(e)

	// repository
	sysRepo := systemRepository.NewSystemRepository(pgClient)
	articleRepo := articleRepository.NewArticleRepository(pgClient)
	categoryRepo := categoryRepository.NewCategoryRepository(pgClient)

	// usecase
	sysUsecase := systemUsecase.NewSystemUsecase(sysRepo)
	articleUC := articleUsecase.NewArticleUsecase(articleRepo)
	categoryUC := categoryUsecase.NewCategoryUsecase(categoryRepo)

	// delivery
	systemDelivery.NewSystemHandler(e, sysUsecase)
	articleDelivery.NewArticleHandler(e, articleUC)
	categoryDelivery.NewCategoryHandler(e, categoryUC)

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
