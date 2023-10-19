package query

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func GetOrderProducts(db *sqlx.DB, orderIds string) (map[string]string, string) {
	//orderMap := make(map[int][]*models.OrderProduct)
	var productIds string
	orderProductMap := make(map[string]string)

	query := `SELECT order_id, product_id, quantity FROM products_orders WHERE order_id IN (` + orderIds + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalf("failed to get orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderID int
		var productID int
		var quantity int
		if err = rows.Scan(&orderID, &productID, &quantity); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		if productIds == "" {
			productIds += fmt.Sprintf("%d", productID)
		} else {
			productIds += fmt.Sprintf(",%d", productID)
		}
		orderProductMap[fmt.Sprintf("%d,%d", orderID, productID)] = fmt.Sprintf("(id=%d)\nЗаказ %d, %d шт", productID, orderID, quantity)

	}

	if err = rows.Err(); err != nil {
		log.Fatalf("error in rows: %v", err)
	}
	return orderProductMap, productIds
}
func GetProductMap(db *sqlx.DB, productIds string) map[int]string {
	productMap := make(map[int]string)
	query := `SELECT * FROM products  WHERE product_id IN (` + productIds + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalf("failed to get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productID int
		var productName string
		if err = rows.Scan(&productID, &productName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		productMap[productID] = productName
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("error in rows: %v", err)
	}
	return productMap
}
func GetProductShelvesMaps(db *sqlx.DB, productIds string) (map[int]int, map[int]string, string) {
	var shelvesIds string
	mainShelfMap := make(map[int]int)
	otherShelfMap := make(map[int]string)
	query := `SELECT * FROM product_shelf  WHERE product_id IN (` + productIds + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalf("failed to get product shelf: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productId int
		var shelfId int
		var mainShelf bool
		if err = rows.Scan(&productId, &shelfId, &mainShelf); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		if shelvesIds == "" {
			shelvesIds += fmt.Sprintf("%d", shelfId)
		} else {
			shelvesIds += fmt.Sprintf(",%d", shelfId)
		}
		if mainShelf {
			mainShelfMap[productId] = shelfId
		} else {
			if _, exists := otherShelfMap[productId]; !exists {
				otherShelfMap[productId] += fmt.Sprintf("%d", shelfId)
			} else {
				otherShelfMap[productId] += fmt.Sprintf(",%d", shelfId)
			}

		}

	}

	if err = rows.Err(); err != nil {
		log.Fatalf("error in rows: %v", err)
	}

	return mainShelfMap, otherShelfMap, shelvesIds
}
func GetShelvesMap(db *sqlx.DB, shelvesIds string) map[int]string {
	shelvesMap := make(map[int]string)
	query := `SELECT * FROM shelves  WHERE shelves.shelf_id IN (` + shelvesIds + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalf("failed to get shelves: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shelveID int
		var shelfName string
		if err = rows.Scan(&shelveID, &shelfName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		shelvesMap[shelveID] = shelfName
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("error in rows: %v", err)
	}
	return shelvesMap
}
