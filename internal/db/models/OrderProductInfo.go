package models

type OrderProductInfo struct {
	OrderID     int    `db:"order_id"`
	ProductName string `db:"product_name"`
	ProductID   int    `db:"product_id"`
	Quantity    int    `db:"quantity"`
	ShelfName   string `db:"shelf_name"`
}
