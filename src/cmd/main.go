package main

import (
	"fmt"
	"sampleServer/src/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handlers.SetRouter(e)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v", r)
		}
	}()

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}

}
