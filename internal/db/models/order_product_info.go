package models

type OrderProductInfo struct {
	OrderID     int
	ProductName string
	ProductID   int
	Quantity    int
	OtherShelf  []string
}
