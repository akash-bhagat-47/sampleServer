package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Details(ctx echo.Context) error {
	data := make(map[string]interface{})

	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "bad request")
	}

	fmt.Println("Received data on /user/details endpoint: ", data)

	return ctx.JSON(http.StatusOK, echo.Map{
		"msg": "details successfully saved",
	})
}
