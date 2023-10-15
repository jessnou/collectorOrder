package query

import (
	"collectorOrder/internal/db/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	get = `
		SELECT s.shelf_name,po.order_id,po.product_id,po.quantity,p.product_name FROM products_orders AS po
		        JOIN orders AS o ON o.order_id = po.order_id
		        JOIN products AS p ON po.product_id = p.product_id
		    	JOIN product_shelf AS ps ON p.product_id = ps.product_id
		        join shelves AS s ON ps.shelf_id = s.shelf_id
    			WHERE po.order_id = $1
    			ORDER BY p.product_id `

	getShelves = `SELECT ps.product_id,p.shelf_name FROM product_shelf AS ps
                    JOIN shelves AS p on ps.shelf_id = p.shelf_id
                    WHERE ps.product_id = $1`
)

func GetOrders(db *sqlx.DB, id int) ([]*models.OrderProductInfo, error) {
	var info []*models.OrderProductInfo
	err := db.Select(&info, get, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	return info, nil
}

func GetShelves(db *sqlx.DB, id int) ([]*models.ProductShelf, error) {
	var info []*models.ProductShelf
	err := db.Select(&info, getShelves, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get shelves: %v", err)
	}
	return info, nil
}
