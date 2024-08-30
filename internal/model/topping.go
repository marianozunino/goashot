package model

import "github.com/marianozunino/goashot/internal/dto"

type Topping struct {
	ID   dto.ToppingType `json:"id"`
	Name string          `json:"name"`
}
