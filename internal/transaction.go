package internal

import (
	"errors"
	"time"

	. "github.com/zoedsoupe/exo/changeset"
)

type Transaction struct {
	ID          int
	Value       int
	CustomerID  int
	Type        [1]string
	Description [10]string
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

func MakeTransaction(ID int, Value map[string]interface{}) (Transaction, error) {
	var t Transaction
	if ID < 1 || ID > 5 {
		return t, errors.New("ID não exitse")
	}

	c := Cast[Transaction](Value)

}
