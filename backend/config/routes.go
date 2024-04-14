package config

import (
	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/api/controllers/incoming_webhooks"
	"github.com/zhall0624/automator3030/api/controllers/workflows"
)

func Routes(e *echo.Echo) {
	e.GET("/incoming-webhooks", incoming_webhooks.Index)
	e.POST("/incoming-webhooks", incoming_webhooks.Create)
	e.POST("/incoming-webhooks/:id", incoming_webhooks.Process)
	e.GET("/workflows/:id", workflows.Show)
	e.GET("/workflows", workflows.Index)
	e.POST("/workflows", workflows.Create)
	e.PATCH("/workflows/:id", workflows.Update)
	e.DELETE("/workflows/:id", workflows.Delete)
}
