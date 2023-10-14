package models

import "time"

type Order struct {
	OrderID      int       `db:"order_id"`
	CustomerName string    `db:"customer_name"`
	OrderDate    time.Time `db:"order_date"`
}
