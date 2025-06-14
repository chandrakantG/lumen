package models

type Order struct {
	ID         uint   `json:"id"`
	CustomerID uint   `json:"customer_id"`
	OrderDate  string `json:"order_date"`
}
