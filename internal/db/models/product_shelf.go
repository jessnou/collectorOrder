package models

type ProductShelf struct {
	ProductId int    `db:"product_id"`
	ShelfName string `db:"shelf_name"`
}
