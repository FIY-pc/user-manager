package router

import (
	"github.com/FIY-pc/user-manager/internal/controller"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	// 根路径
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Please use api")
	})
	// tokens
	e.GET("/tokens", controller.Login())
	// users
	e.POST("/users/register", controller.Register())
	e.GET("/users", controller.GetUser())
	e.POST("/users", controller.CreateUser())
	e.PUT("/users", controller.UpdateUser())
	e.DELETE("/users", controller.DeleteUser())
}
