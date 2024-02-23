package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	// e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/clientes/:id/transacoes", handleTransaction)
	e.Logger.Fatal(e.Start(*port))
}

func handleTransaction(c echo.Context) error {
	param := c.Param("id")

	var req TransactionRequest
	err := (&echo.DefaultBinder{}).BindBody(c, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "bind req")
	}

	customerID, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "atoi")
	}

	var response TransactionResponse
	cc, err := MakeTransaction(customerID, req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	response.Saldo = cc.Balance
	response.Limite = cc.MaxLimit

	body, err := json.Marshal(response)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "marshal response")
	}

	fmt.Println(string(body))
	return c.JSON(http.StatusOK, body)
}
