package query

import (
	"collectorOrder/internal/db/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetOrders(db *sqlx.DB, id int) ([]*models.OrderProduct, error) {
	var info []*models.OrderProduct
	err := db.Select(&info, `SELECT * FROM products_orders WHERE order_id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	return info, nil
}

func GetProductShelf(db *sqlx.DB, id int) ([]*models.ProductShelf, error) {
	var info []*models.ProductShelf
	err := db.Select(&info, `SELECT * FROM products AS p 
         JOIN product_shelf AS ps on p.product_id = ps.product_id
         WHERE p.product_id = $1 `, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	return info, nil
}

func GetShelf(db *sqlx.DB, id int) (*models.Shelves, error) {
	var info []*models.Shelves
	err := db.Select(&info, `SELECT * FROM shelves WHERE shelf_id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get shelves: %v", err)
	}
	return info[0], nil
}
