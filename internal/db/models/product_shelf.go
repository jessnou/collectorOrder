package models

type ProductShelf struct {
	ProductId int  `db:"product_id"`
	ShelfId   int  `db:"shelf_id"`
	MainShelf bool `db:"main_shelf"`
}
