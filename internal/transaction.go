package internal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	. "github.com/zoedsoupe/exo/changeset"
)

type Transaction struct {
	ID          int
	Value       int
	CustomerID  int
	Type        string
	Description string
	CreatedAt   time.Time
}

var transactions []Transaction

var client = [5]Client{
	{1, "A", 1000 * 100, 0},
	{2, "B", 800 * 100, 0},
	{3, "C", 10000 * 100, 0},
	{4, "D", 100000 * 100, 0},
	{5, "E", 5000 * 100, 0},
}

func MakeTransaction(ID int, Value map[string]interface{}) (Client, error) {
	var t Transaction
	var client Client
	if ID < 1 || ID > 5 {
		return client, echo.NewHTTPError(http.StatusNotFound, nil)
	}

	c := Cast[Transaction](Value).
		ValidateChange("Type", InclusionValidator{Allowed: []interface{}{"c", "d"}}).
		ValidateChange("Description", LengthValidator{Min: 1, Max: 10})

	t, err := ApplyNew[Transaction](c)
	if err != nil {
		return client, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	transactions = append(transactions, t)

	return client, nil
}

func processTransaction(client *Client, trx Transaction) error {
	if trx.Type == "c" {
		client.Balance += trx.Value
		return nil
	}

	if !validTransaction(*client, trx.Value) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, nil)
	}

	client.Balance -= trx.Value
	return nil
}

func validTransaction(cilent Client, value int) bool {
	return true
}
