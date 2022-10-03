package service

import (
	"errors"

	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/mapper"
	storage "github.com/marianozunino/goashot/internal/storage/json"
)

type OrderService interface {
	GetOrders() []*dto.Order
	GetOrder(id int) (*dto.Order, error)
	AddOrder(order *dto.Order)
	UpdateOrder(order *dto.Order)
	DeleteOrder(id int)
}

func registerOrderService(repo storage.Repository) OrderService {
	return &orderService{
		repo: repo,
	}
}

type orderService struct {
	repo storage.Repository
}

// AddOrder implements OrderService
func (s *orderService) AddOrder(order *dto.Order) {
	s.repo.AddOrder(mapper.OrderDtoToOrderModel(order))
}

// DeleteOrder implements OrderService
func (s *orderService) DeleteOrder(id int) {
	s.repo.DeleteOrder(id)
}

// GetOrder implements OrderService
func (o *orderService) GetOrder(id int) (*dto.Order, error) {
	order := o.repo.GetOrder(id)
	if order == nil {
		return nil, errors.New("order not found")
	}
	return mapper.OrderModelToOrderDto(order), nil
}

// GetOrders implements OrderService
func (s *orderService) GetOrders() []*dto.Order {
	orders := s.repo.GetOrders()
	orderDtos := mapper.OrderModelsToOrderDtos(orders)
	return orderDtos
}

// UpdateOrder implements OrderService
func (s *orderService) UpdateOrder(order *dto.Order) {
	s.repo.UpdateOrder(mapper.OrderDtoToOrderModel(order))
}
