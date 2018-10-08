package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func (e *EchoSkeleton) helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
