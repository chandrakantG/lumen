package models

type OrderItem struct {
	ID        uint `json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
