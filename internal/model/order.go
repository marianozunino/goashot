package model

import "github.com/marianozunino/goashot/internal/dto"

type Order struct {
	ID        int              `json:"id"`
	OrderType dto.ShawarmaType `json:"orderType"`
	Toppings  []Topping        `json:"toppings"`
	User      string           `json:"user"`
	IsActive  bool             `json:"isActive"`
}
