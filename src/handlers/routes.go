package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo) {

	// Health check route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"response": "Health check ok. server is up!"})
	})

	v1API := e.Group("/v1")
	v1API.POST("/user/details", Details)
}
