package main

import (
	"flag"
	"net/http"
	"strconv"

	. "github.com/cciuenf/rinha/internal"

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

	e.GET("/clientes/:id/transacoes", handleTransaction)
	e.Logger.Fatal(e.Start(*port))
}

func handleTransaction(c echo.Context) error {
	param := c.Param("id")

	var attrs map[string]interface{}
	err := (&echo.DefaultBinder{}).BindBody(c, &attrs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	customerID, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	t, err := MakeTransaction(customerID, attrs)

	return nil
}
