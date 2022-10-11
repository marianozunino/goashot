package json

import (
	"sync"

	"github.com/marianozunino/goashot/internal/model"
)

type Repository interface {
	GetOrders() []*model.Order
	GetOrder(id int) *model.Order
	AddOrder(order *model.Order)
	UpdateOrder(order *model.Order)
	DeleteOrder(id int)
}

type repository struct {
	sync.Mutex
	Database
}

// assume that the repository implements the Repository interface
var _ Repository = (*repository)(nil)

func registerRepository(db Database) Repository {
	r := &repository{
		sync.Mutex{},
		db,
	}
	return r
}

func (r *repository) GetOrders() []*model.Order {
	return r.orders
}

func (r *repository) GetOrder(id int) *model.Order {
	for _, order := range r.orders {
		if order.ID == id {
			return order
		}
	}
	return nil
}

func (r *repository) AddOrder(order *model.Order) {
	r.Lock()
	defer r.persistOrders()
	defer r.Unlock()

	order.ID = r.getNewID()
	r.orders = append(r.orders, order)
}

func (r *repository) UpdateOrder(order *model.Order) {
	r.Lock()
	defer r.persistOrders()
	defer r.Unlock()

	for i, o := range r.orders {
		if o.ID == order.ID {
			r.orders[i] = order
		}
	}
}

func (r *repository) DeleteOrder(id int) {
	r.Lock()
	defer r.persistOrders()
	defer r.Unlock()

	orders := make([]*model.Order, 0)
	for i, order := range r.orders {
		if order.ID == id {
			orders = append(r.orders[:i], r.orders[i+1:]...)
		}
	}
	r.orders = orders
}
