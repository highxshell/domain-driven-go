// Package aggregates holds aggregates that combines many entities into a full object
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/highxshell/tavern"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *tavern.Person
	// a customer can hold many products
	product []*tavern.Item
	// a customer can perform many transactions
	transaction []tavern.Transaction
}

// New is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func New(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:      person,
		product:     make([]*tavern.Item, 0),
		transaction: make([]tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}
