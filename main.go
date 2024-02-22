package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := flag.String("port", ":4000", "port that server will listen")

	e := echo.New()

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", handleBasic)
	e.Logger.Fatal(e.Start(*port))
}

func handleBasic(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
