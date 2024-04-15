package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/models"
)

func GetStep(c echo.Context) error {
	var step models.Step

	err := c.Bind(&step)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	res := models.DB.Find(&step)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "step not found"})
	}

	return c.JSON(http.StatusOK, &step)
}

func CreateStep(c echo.Context) error {
	var step models.Step

	err := c.Bind(&step)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	res := models.DB.Create(&step)
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	return c.JSON(http.StatusCreated, &step)
}

func UpdateStep(c echo.Context) error {
	var step models.Step

	err := c.Bind(&step)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	res := models.DB.Save(&step)
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	return c.JSON(http.StatusOK, &step)
}
