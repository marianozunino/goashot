package service

import (
	"errors"

	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/mapper"
	"github.com/marianozunino/goashot/internal/model"
)

type Repository interface {
	GetOrders() []*model.Order
	GetOrder(id int) *model.Order
	AddOrder(order *model.Order) (*model.Order, error)
	UpdateOrder(order *model.Order) (*model.Order, error)
	DeleteOrder(id int)
}

func NewOrderService(repo Repository) OrderService {
	return OrderService{
		repo: repo,
	}
}

type OrderService struct {
	repo Repository
}

// AddOrder implements OrderService
func (s *OrderService) AddOrder(order *dto.Order) (*dto.Order, error) {
	createdOrder, err := s.repo.AddOrder(mapper.OrderDtoToOrderModel(order))
	if err != nil {
		return nil, err
	}
	return mapper.OrderModelToOrderDto(createdOrder), nil
}

// DeleteOrder implements OrderService
func (s *OrderService) DeleteOrder(id int) {
	s.repo.DeleteOrder(id)
}

// GetOrder implements OrderService
func (o *OrderService) GetOrder(id int) (*dto.Order, error) {
	order := o.repo.GetOrder(id)
	if order == nil {
		return nil, errors.New("order not found")
	}

	return mapper.OrderModelToOrderDto(order), nil
}

// GetOrders implements OrderService
func (s *OrderService) GetOrders() []*dto.Order {
	orders := s.repo.GetOrders()
	orderDtos := mapper.OrderModelsToOrderDtos(orders)
	return orderDtos
}

// UpdateOrder implements OrderService
func (s *OrderService) UpdateOrder(order *dto.Order) (*dto.Order, error) {

	updatedOrder, err := s.repo.UpdateOrder(mapper.OrderDtoToOrderModel(order))
	if err != nil {
		return nil, err
	}
	return mapper.OrderModelToOrderDto(updatedOrder), nil

}
