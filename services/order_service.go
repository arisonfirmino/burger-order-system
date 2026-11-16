package services

import (
	"fmt"
	"math/rand"

	"burger-order-system/models"
)

func generateID() string {
	return fmt.Sprintf("#%05d", rand.Intn(100000))
}

var orders []models.Order

func CreateOrder(client string) models.Order {
	return models.Order{
		ID:       generateID(),
		Client:   client,
		Status:   "Aberto",
		Products: []*models.OrderItem{},
	}
}

func AddProduct(order *models.Order, product *models.Product, quantity int) error {
	if order == nil {
		return fmt.Errorf("Pedido não encontrado.")
	}

	if product == nil {
		return fmt.Errorf("Produto não encontrado.")
	}

	if quantity <= 0 {
		return fmt.Errorf("A quantidade deve ser maior que zero.")
	}

	item := &models.OrderItem{
		Product:  product,
		Quantity: quantity,
	}

	order.Products = append(order.Products, item)

	return nil
}

func CalculateTotal(order *models.Order) float64 {
	total := 0.0

	for _, product := range order.Products {
		total += product.Product.Price * float64(product.Quantity)
	}

	return total
}

func SaveOrder(order models.Order) {
	orders = append(orders, order)
}

func GetOrders() []models.Order {
	return orders
}
