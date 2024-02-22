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

type TransactionResponse struct {
	Saldo int `json:"saldo"`
	Limite int `json:"limite"`
}

func handleTransaction(c echo.Context) error {
	fmt.Println("Ola")
	param := c.Param("id")

	var attrs map[string]interface{}
	err := (&echo.DefaultBinder{}).BindBody(c, &attrs)
	fmt.Println("Ola2")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	customerID, err := strconv.Atoi(param)
	fmt.Println("Ola3")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	var response TransactionResponse
	cc, err := MakeTransaction(customerID, attrs)
	response.Saldo = cc.Balance
	response.Limite = cc.MaxLimit

	body, err := json.Marshal(response)
	fmt.Println("Ola4")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	fmt.Println("Ola5")
	fmt.Println(string(body))
	return c.JSON(http.StatusOK, string(body))
}
