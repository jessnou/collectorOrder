package query

import (
	"collectorOrder/internal/db/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	get = `
	SELECT s.shelf_name,o.order_id, p.product_name, p.product_id, op.quantity 
        FROM orders AS o
        JOIN order_products AS op ON o.order_id = op.order_id
        JOIN products AS p ON op.product_id = p.product_id
        JOIN shelves as s ON p.main_shelf_id = s.shelf_id
        WHERE o.order_id in ($1)
        ORDER BY s.shelf_id `
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
