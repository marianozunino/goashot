package mapper

import (
	"log"

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
	for _, toppingType := range orderDto.Toppings {
		topping, err := dto.GetTopping(toppingType)
		if err != nil {
			log.Println(err)
			continue
		}
		toppings = append(toppings, model.Topping{
			ID:   topping.ID,
			Name: topping.Name,
		})
	}

	result := &model.Order{
		ID:        orderDto.ID,
		OrderType: orderDto.OrderType,
		Toppings:  toppings,
		User:      orderDto.User,
		IsActive:  orderDto.IsActive,
	}
	return result
}
