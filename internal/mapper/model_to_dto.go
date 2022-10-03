package mapper

import (
	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/model"
)

func OrderDtosToOrderModels(ordersDto []*dto.Order) []*model.Order {
	var orders []*model.Order
	for _, order := range ordersDto {
		orders = append(orders, OrderDtoToOrderModel(order))
	}
	return orders
}
func OrderDtoToOrderModel(orderDto *dto.Order) *model.Order {
	var toppings []model.Topping
	for _, topping := range orderDto.Toppings {
		toppingName, err := dto.Toppings.GetToppingName(topping)
		if err != nil {
			continue
		}
		toppings = append(toppings, model.Topping{
			ID:   string(topping),
			Name: string(toppingName),
		})
	}
	return &model.Order{
		ID:        orderDto.ID,
		OrderType: orderDto.OrderType,
		Toppings:  toppings,
		User:      orderDto.User,
		IsActive:  orderDto.IsActive,
	}
}
