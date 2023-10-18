package models

type Product struct {
	ProductID   int    `db:"product_id"`
	ProductName string `db:"product_name"`
}
