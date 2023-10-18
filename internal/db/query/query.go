package query

import (
	"collectorOrder/internal/db/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetOrderProducts(db *sqlx.DB, ids string) (map[int][]*models.OrderProduct, error) {
	orderMap := make(map[int][]*models.OrderProduct)

	query := `SELECT order_id, product_id, quantity FROM products_orders WHERE order_id IN (` + ids + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.OrderProduct
		if err = rows.Scan(&order.OrderID, &order.ProductID, &order.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		if _, exists := orderMap[order.ProductID]; !exists {
			orderMap[order.ProductID] = []*models.OrderProduct{}
		}
		orderMap[order.ProductID] = append(orderMap[order.ProductID], &order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %v", err)
	}

	return orderMap, nil
}

func GetProductShelf(db *sqlx.DB, ids string) (map[int][]*models.ProductShelf, error) {
	productShelfMap := make(map[int][]*models.ProductShelf)
	query := `SELECT * FROM product_shelf  WHERE product_id IN (` + ids + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get product shelf: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productShelf models.ProductShelf
		if err = rows.Scan(&productShelf.ProductId, &productShelf.ShelfId, &productShelf.MainShelf); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		if _, exists := productShelfMap[productShelf.ShelfId]; !exists {
			productShelfMap[productShelf.ShelfId] = []*models.ProductShelf{}
		}
		productShelfMap[productShelf.ShelfId] = append(productShelfMap[productShelf.ProductId], &productShelf)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %v", err)
	}

	return productShelfMap, nil
}
func GetProduct(db *sqlx.DB, ids string) (map[int]*models.Product, error) {
	productMap := make(map[int]*models.Product)
	query := `SELECT * FROM products  WHERE product_id IN (` + ids + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err = rows.Scan(&product.ProductID, &product.ProductName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		productMap[product.ProductID] = &product
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %v", err)
	}

	return productMap, nil
}

func GetShelves(db *sqlx.DB, ids string) (map[int]*models.Shelves, error) {
	shelvesMap := make(map[int]*models.Shelves)
	query := `SELECT * FROM shelves  WHERE shelves.shelf_id IN (` + ids + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get shelves: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shelf models.Shelves
		if err = rows.Scan(&shelf.ShelveID, &shelf.ShelfName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		shelvesMap[shelf.ShelveID] = &shelf
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %v", err)
	}

	return shelvesMap, nil
}
