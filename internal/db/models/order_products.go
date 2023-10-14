package models

type OrderProduct struct {
	OrderID   int `db:"order_id"`
	ProductID int `db:"product_id"`
	Quantity  int `db:"quantity"`
}
