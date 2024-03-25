package config

import (
	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/api/controllers/incoming_webhooks"
)

func Routes(e *echo.Echo) {
	e.GET("/incoming-webhooks", incoming_webhooks.Index)
	e.POST("/incoming-webhooks", incoming_webhooks.Create)
	e.POST("/incoming-webhooks/:id", incoming_webhooks.Process)
}
