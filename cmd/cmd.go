package main

import (
	"fmt"
	"github.com/FIY-pc/user-manager/internal/config"
	"github.com/FIY-pc/user-manager/internal/model"
	"github.com/FIY-pc/user-manager/internal/router"
	"github.com/FIY-pc/user-manager/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(util.JWTAuthMiddleware())
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	url := fmt.Sprintf("%s:%s", config.Config.Server.Host, config.Config.Server.Port)
	e.Logger.Fatal(e.Start(url))
}
