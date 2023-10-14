package models

type Shelves struct {
	ShelveID  int    `db:"shelf_id"`
	ShelfName string `db:"shelf_name"`
}
