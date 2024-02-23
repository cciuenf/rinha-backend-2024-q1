package internal

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID          int
	Value       int
	CustomerID  int
	Type        string
	Description string
	CreatedAt   time.Time
}

type TransactionRequest struct {
	Description string `json:"descricao"`
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
}

type TransactionResponse struct {
	Saldo  int `json:"saldo"`
	Limite int `json:"limite"`
}

var transactions []Transaction

var clients = [5]Client{
	{1, "A", 1000 * 100, 0},
	{2, "B", 800 * 100, 0},
	{3, "C", 10000 * 100, 0},
	{4, "D", 100000 * 100, 0},
	{5, "E", 5000 * 100, 0},
}

func MakeTransaction(ID int, req TransactionRequest) (Client, error) {
	var client Client
	if ID < 1 || ID > 5 {
		return client, echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	client = clients[ID - 1]

	t, err := transactionRequestToTransaction(ID, req)
	if err != nil {
		fmt.Println(err.Error())
		return client, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	transactions = append(transactions, t)
	if err := processTransaction(&client, t); err != nil {
		return client, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	clients[ID - 1] = client
	return client, nil
}

func transactionRequestToTransaction(ID int, req TransactionRequest) (Transaction, error) {
	var t Transaction

	if len(req.Description) > 10 {
		return t, echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid desc")
	}

	if (req.Type != "c") && (req.Type != "d") {
		return t, echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid type")
	}

	t.ID = rand.Int()
	t.CustomerID = ID
	t.CreatedAt = time.Now()
	t.Type = req.Type
	t.Description = req.Description
	t.Value = req.Value

	return t, nil
}

func processTransaction(client *Client, trx Transaction) error {
	if trx.Type == "c" {
		client.Balance += trx.Value
		return nil
	}

	if invalidTransaction(*client, trx.Value) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "can't process trx")
	}

	client.Balance -= trx.Value
	return nil
}

func invalidTransaction(client Client, value int) bool {
	return client.Balance-value < (-1 * client.MaxLimit)
}
