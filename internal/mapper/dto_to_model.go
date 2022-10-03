package mapper

import (
	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/model"
)

func OrderModelToOrderDto(orderModel *model.Order) *dto.Order {
	var toppings []dto.ToppingKey
	for _, topping := range orderModel.Toppings {
		toppings = append(toppings, dto.ToppingKey(topping.ID))
	}
	return &dto.Order{
		ID:        orderModel.ID,
		OrderType: orderModel.OrderType,
		Toppings:  toppings,
		User:      orderModel.User,
		IsActive:  orderModel.IsActive,
	}
}

func OrderModelsToOrderDtos(ordersModel []*model.Order) []*dto.Order {
	var orders []*dto.Order
	for _, order := range ordersModel {
		orders = append(orders, OrderModelToOrderDto(order))
	}
	return orders
}
