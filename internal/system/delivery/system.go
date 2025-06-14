package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ishtiaqhimel/news-portal/cms/internal/response"
	"github.com/ishtiaqhimel/news-portal/cms/internal/system/usecase"
)

// SystemHandler represents the httpHandler for system
type SystemHandler struct {
	Usecase usecase.SystemUsecase
}

// NewSystemHandler will initialize the system endpoint
func NewSystemHandler(e *echo.Echo, us usecase.SystemUsecase) {
	handler := &SystemHandler{
		Usecase: us,
	}

	e.GET("/", handler.Root)
	e.GET("/h34l7h", handler.Health)
	e.GET("/api/v1/server-time", handler.ServerTime)
}

// Root will let you see what you can slash 🐲
func (sh *SystemHandler) Root(c echo.Context) error {
	return c.JSON(response.RespondSuccess("Hello there, I'm News Portal CMS!!!", nil))
}

// Health will let you know the heart beats ❤️
func (sh *SystemHandler) Health(c echo.Context) error {
	resp, err := sh.Usecase.GetHealth()
	if err != nil {
		return c.JSON(response.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// ServerTime will let you know the current time on server ⏰
func (sh *SystemHandler) ServerTime(c echo.Context) error {
	resp := sh.Usecase.GetTime()
	return c.JSON(http.StatusOK, resp)
}
