package services

import "burger-order-system/models"

func GetProducts() []models.Product {
	products := []models.Product{
		{
			ID:          "L001",
			Name:        "Classic Burger",
			Description: "Pão, hambúrguer artesanal, queijo cheddar e molho especial.",
			Price:       24.90,
		},
		{
			ID:          "L002",
			Name:        "Bacon Burger",
			Description: "Pão, hambúrguer artesanal, queijo, bacon crocante e molho barbecue.",
			Price:       29.90,
		},
		{
			ID:          "B001",
			Name:        "Refrigerante",
			Description: "Lata de refrigerante 350 ml.",
			Price:       6.0,
		},
		{
			ID:          "B002",
			Name:        "Suco Natural",
			Description: "Suco natural de laranja 300 ml.",
			Price:       8.0,
		},
		{
			ID:          "P001",
			Name:        "Batata Frita",
			Description: "Porção de batatas fritas crocantes.",
			Price:       14.90,
		},
	}

	return products
}
