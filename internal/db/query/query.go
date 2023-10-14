package query

import (
	"collectorOrder/internal/db/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	getOrder = `SELECT * FROM orders WHERE order_id = $1`
	getJoin  = `
	SELECT o.order_id, p.product_name, p.product_id, op.quantity, s.shelf_name
        FROM orders AS o
        JOIN order_products AS op ON o.order_id = op.order_id
        JOIN products AS p ON op.product_id = p.product_id
        JOIN shelves as s ON p.main_shelf_id = s.shelf_id
        WHERE o.order_id = $1
        ORDER BY p.main_shelf_id, p.product_name`
	//getProducts = `SELECT * FROM order_products WHERE order_id = $1`
	//getProduct  = `SELECT * FROM products WHERE product_id = $1`
)

func Get(db *sqlx.DB, id int) ([]*models.OrderProductInfo, error) {
	var info []*models.OrderProductInfo
	err := db.Select(&info, getJoin, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get timezone: %v", err)
	}
	return info, nil
}

//func GetOrder(db *sqlx.DB, id int) (*models.Order, error) {
//	var order []*models.Order
//	err := db.Select(&order, getOrder, id)
//	if err == sql.ErrNoRows {
//		return nil, nil
//	} else if err != nil {
//		return nil, fmt.Errorf("failed to get timezone: %v", err)
//	}
//	return order[0], nil
//}
//
//func GetProducts(db *sqlx.DB, id int) ([]*models.OrderProduct, error) {
//	var orderProduct []*models.OrderProduct
//	err := db.Select(&orderProduct, getProducts, id)
//	if err == sql.ErrNoRows {
//		return nil, nil
//	} else if err != nil {
//		return nil, fmt.Errorf("failed to get timezone: %v", err)
//	}
//	return orderProduct, nil
//}
//
//func GetProduct(db *sqlx.DB, id int) ([]*models.Product, error) {
//	var product []*models.Product
//	err := db.Select(&product, getProduct, id)
//	if err == sql.ErrNoRows {
//		return nil, nil
//	} else if err != nil {
//		return nil, fmt.Errorf("failed to get timezone: %v", err)
//	}
//	return product, nil
//}
