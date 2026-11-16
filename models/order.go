package models

type Order struct {
	ID       string
	Client   string
	Products []*OrderItem
	Status   string
}
