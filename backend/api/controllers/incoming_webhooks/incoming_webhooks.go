package incoming_webhooks

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/models"
)

func Index(c echo.Context) error {
	var webhooks []models.IncomingWebHook
	models.DB.Find(&webhooks)
	return c.JSON(http.StatusOK, webhooks)
}

func Create(c echo.Context) error {
	var webhook models.IncomingWebHook
	c.Bind(&webhook)
	result := models.DB.Create(&webhook)

	if result.Error != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "Unprocessable Entry",
		})
	}
	return c.JSON(http.StatusCreated, webhook)
}

func Process(c echo.Context) error {
	id, conv_err := strconv.Atoi(c.Param("id"))
	if conv_err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "Unprocessable Entry",
		})
	}
	payload, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "Unprocessable Entry",
		})
	}

	webhook_payload := models.IncomingWebHookPayload{IncomingWebHookID: id, Payload: string(payload[:])}
	result := models.DB.Create(&webhook_payload)

	if result.Error != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "Unprocessable Entry",
		})
	}
	return c.String(http.StatusCreated, "")
}
