package router

import (
	"redigo-example/handler"

	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, h *handler.UserHandler) {
	e.GET("/withredis", h.GetAll)
	e.GET("/withoutredis", h.GetAllWithoutRedis)
	e.GET("/:id", h.GetDetail)
	e.POST("/", h.Create)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Delete)
}
