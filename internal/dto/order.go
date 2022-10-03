package dto

type Order struct {
	ID        int          `json:"id" form:"id"`
	OrderType ShawarmaType `json:"orderType" form:"orderType"`
	Toppings  []ToppingKey `json:"toppings" form:"toppings"`
	User      string       `json:"user" form:"user"`
	IsActive  bool         `json:"isActive" form:"isActive"`
}
