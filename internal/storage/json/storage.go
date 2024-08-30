package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marianozunino/goashot/internal/model"
)

func NewDB() *Database {
	db := &Database{}
	db.loadOrders()
	return db
}

type Database struct {
	orders []*model.Order
}

func (db *Database) getNewID() int {
	maxID := 0
	for _, order := range db.orders {
		if order.ID > maxID {
			maxID = order.ID
		}
	}
	return maxID + 1
}

// If any error occurs, the function returns an empty slice
func (db *Database) loadOrders() {
	orders := make([]*model.Order, 0)

	file, err := os.ReadFile("orders.json")
	if err != nil {
		orders = make([]*model.Order, 0)
	}

	err = json.Unmarshal(file, &orders)
	if err != nil {
		orders = make([]*model.Order, 0)
	}

	fmt.Println("---------------   DB LOADED   ---------------")
	fmt.Printf("Loaded %d orders from file\n", len(orders))
	db.orders = orders
}

func (db *Database) persistOrders() {
	file, _ := json.MarshalIndent(db.orders, "", " ")
	_ = os.WriteFile("orders.json", file, 0644)
}
