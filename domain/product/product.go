package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/highxshell/tavern"
)

var (
	ErrMissingValue = errors.New("missing values")
)

type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

func New(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: name,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
