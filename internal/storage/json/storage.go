package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marianozunino/goashot/internal/model"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(registerRepository),
	fx.Provide(registerDB),
)

type Database = *database

func registerDB() Database {
	db := &database{}
	db.loadOrders()
	return db
}

type database struct {
	orders []*model.Order
}

func (db *database) getNewID() int {
	maxID := 0
	for _, order := range db.orders {
		if order.ID > maxID {
			maxID = order.ID
		}
	}
	return maxID + 1
}

// If any error occurs, the function returns an empty slice
func (db *database) loadOrders() {
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

func (db *database) persistOrders() {
	file, _ := json.MarshalIndent(db.orders, "", " ")
	_ = os.WriteFile("orders.json", file, 0644)
}
