package main

import (
	"github.com/google/uuid"
	"github.com/highxshell/tavern/domain/product"
	"github.com/highxshell/tavern/services/order"
	"github.com/highxshell/tavern/services/tavern"
)

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		//order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}
	uid, err := os.AddCustomer("Artem")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	if err := tavern.Order(uid, order); err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.New("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.New("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.New("Wine", "Nasty Drink", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
