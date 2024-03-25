package config

import (
	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/models"
)

func Setup() {
	models.ConnectDatabase()
	e := echo.New()
	Middleware(e)
	Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
