package main

import (
	"fmt"

	"github.com/labstack/echo"
)

type EchoSkeleton struct {
	echoServer *echo.Echo
}

func Init() *EchoSkeleton {
	return &EchoSkeleton{
		echoServer: echo.New(),
	}
}

func (e *EchoSkeleton) setupHandlers() {
	e.echoServer.GET("/", e.helloWorldHandler)
}

func (e *EchoSkeleton) Start(port int) {
	e.setupHandlers()
	e.echoServer.Logger.Fatal(e.echoServer.Start(fmt.Sprintf(":%d", port)))
}

func main() {
	echoSkeleton := Init()
	echoSkeleton.Start(8080)
}
