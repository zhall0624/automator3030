package config

import (
	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/api/controllers"
)

func Routes(e *echo.Echo) {
	e.GET("/incoming-webhooks", controllers.GetIncomingWebhooks)
	e.POST("/incoming-webhooks/:id", controllers.CreateIncomingWebhook)
	e.POST("/incoming-webhooks", controllers.ProcessIncomingWebhook)

	e.GET("/workflows/:id", controllers.GetWorkflow)
	e.PATCH("/workflows/:id", controllers.UpdateWorkflow)
	e.GET("/workflows", controllers.GetAllWorkflows)
	e.POST("/workflows", controllers.CreateWorkflow)
	e.DELETE("/workflows/:id", controllers.DeleteWorkflow)
}
