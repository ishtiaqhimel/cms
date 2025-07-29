package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
)

const EchoLogFormat = "time: ${time_rfc3339_nano} || ${method}: ${uri} || u_agent: ${user_agent} || status: ${status} || latency: ${latency_human} \n"

func Attach(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: EchoLogFormat}))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit(config.Get().App.RequestBodyLimit))
	e.Pre(middleware.RemoveTrailingSlash())
}
