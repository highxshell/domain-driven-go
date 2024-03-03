package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/highxshell/tavern/domain/customer"
	"github.com/highxshell/tavern/domain/product"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.New("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := product.New("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := product.New("Wine", "Nasty Drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderSeervice(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	cust, err := customer.New("Artem")
	if err != nil {
		t.Error(err)
	}
	if err = os.customers.Add(cust); err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
