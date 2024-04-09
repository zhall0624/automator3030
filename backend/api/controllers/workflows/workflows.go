package workflows

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhall0624/automator3030/models"
)

func Index(c echo.Context) error {
	var workflows []models.Workflow
	models.DB.Find(&workflows)
	return c.JSON(http.StatusOK, &workflows)
}

func Show(c echo.Context) error {
	var workflow models.Workflow

	if err := c.Bind(&workflow); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	if err := models.DB.Find(&workflow); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Workflow not found"})
	}

	return c.JSON(http.StatusOK, &workflow)
}

func Create(c echo.Context) error {
	var workflow models.Workflow
	if err := c.Bind(&workflow); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	models.DB.Create(&workflow)

	return c.JSON(http.StatusCreated, &workflow)
}

func Update(c echo.Context) error {
	var workflow models.Workflow
	if err := c.Bind(&workflow); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	if err := models.DB.Save(&workflow); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Workflow could not be updated"})
	}
	return c.JSON(http.StatusOK, &workflow)
}

func Delete(c echo.Context) error {
	var workflow models.Workflow
	if err := c.Bind(workflow); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	if err := models.DB.Find(&workflow); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Workflow not found"})
	}
	if err := models.DB.Delete(&workflow); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete Workflow"})
	}
	return c.JSON(http.StatusOK, map[string]string{"error": "Workflow not found"})
}
