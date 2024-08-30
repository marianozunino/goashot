package json

import (
	"sync"

	"github.com/marianozunino/goashot/internal/model"
	"github.com/marianozunino/goashot/internal/service"
)

type Repository struct {
	sync.Mutex
	db *Database
}

// assume that the repository implements the Repository interface
var _ service.Repository = (*Repository)(nil)

func NewRepository(db *Database) *Repository {
	r := &Repository{
		sync.Mutex{},
		db,
	}
	return r
}

func (r *Repository) GetOrders() []*model.Order {
	return r.db.orders
}

func (r *Repository) GetOrder(id int) *model.Order {
	for _, order := range r.db.orders {
		if order.ID == id {
			return order
		}
	}
	return nil
}

func (r *Repository) AddOrder(order *model.Order) (*model.Order, error) {
	r.Lock()
	defer r.db.persistOrders()
	defer r.Unlock()

	order.ID = r.db.getNewID()
	r.db.orders = append(r.db.orders, order)

	return order, nil
}

func (r *Repository) UpdateOrder(order *model.Order) (*model.Order, error) {
	r.Lock()
	defer r.db.persistOrders()
	defer r.Unlock()

	for i, o := range r.db.orders {
		if o.ID == order.ID {
			r.db.orders[i] = order
		}
	}

	return order, nil
}

func (r *Repository) DeleteOrder(id int) {
	r.Lock()
	defer r.db.persistOrders()
	defer r.Unlock()

	orders := make([]*model.Order, 0)
	for i, order := range r.db.orders {
		if order.ID == id {
			orders = append(r.db.orders[:i], r.db.orders[i+1:]...)
		}
	}
	r.db.orders = orders
}
