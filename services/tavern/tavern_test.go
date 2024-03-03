package tavern

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/highxshell/tavern/domain/product"
	"github.com/highxshell/tavern/services/order"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	uid, err := os.AddCustomer("Artem")
	if err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	if err = tavern.Order(uid, order); err != nil {
		t.Fatal(err)
	}
}

func Test_MongoTavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}
	uid, err := os.AddCustomer("Artem")
	if err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}

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
