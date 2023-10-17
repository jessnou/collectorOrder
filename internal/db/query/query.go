package query

import (
	"collectorOrder/internal/db/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetOrders(db *sqlx.DB, ids string) ([]*models.OrderProduct, error) {
	var info []*models.OrderProduct

	err := db.Select(&info, fmt.Sprintf("SELECT * FROM products_orders WHERE order_id IN  %s ", ids))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	return info, nil
}

func GetProductShelf(db *sqlx.DB, ids string) ([]*models.ProductShelf, error) {
	var info []*models.ProductShelf
	err := db.Select(&info, fmt.Sprintf(`SELECT * FROM products AS p
         JOIN product_shelf AS ps on p.product_id = ps.product_id
         WHERE p.product_id IN  %s `, ids))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	return info, nil
}

func GetShelves(db *sqlx.DB, ids string) ([]*models.Shelves, error) {
	var info []*models.Shelves
	err := db.Select(&info, fmt.Sprintf(`SELECT * FROM shelves WHERE shelf_id IN %s`, ids))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get shelves: %v", err)
	}
	return info, nil
}
