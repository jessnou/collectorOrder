package query

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

func GetOrderProducts(db *sqlx.DB, orderIds string) error {
	//orderMap := make(map[int][]*models.OrderProduct)
	var productIds string

	orderProductMap := make(map[int]string)
	//orderProductInfo := make(map[string]string)

	query := `SELECT order_id, product_id, quantity FROM products_orders WHERE order_id IN (` + orderIds + `)`

	rows, err := db.Queryx(query)
	if err != nil {
		return fmt.Errorf("failed to get orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var orderID int
		var productID int
		var quantity int
		if err = rows.Scan(&orderID, &productID, &quantity); err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}
		if productIds == "" {
			productIds += fmt.Sprintf("%d", productID)
		} else {
			productIds += fmt.Sprintf(",%d", productID)
		}
		orderProductMap[productID] += fmt.Sprintf("(id=%d) Заказ %d, %d шт", productID, orderID, quantity)

	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error in rows: %v", err)
	}

	productMap := make(map[int]string)
	query = `SELECT * FROM products  WHERE product_id IN (` + productIds + `)`

	rows, err = db.Queryx(query)
	if err != nil {
		return fmt.Errorf("failed to get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productID int
		var productName string
		if err = rows.Scan(&productID, &productName); err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}
		productMap[productID] = productName
	}
	var shelvesIds string
	mainShelfMap := make(map[int]int)
	otherShelfMap := make(map[int]string)
	query = `SELECT * FROM product_shelf  WHERE product_id IN (` + productIds + `)`

	rows, err = db.Queryx(query)
	if err != nil {
		return fmt.Errorf("failed to get product shelf: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productId int
		var shelfId int
		var mainShelf bool
		if err = rows.Scan(&productId, &shelfId, &mainShelf); err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}
		if shelvesIds == "" {
			shelvesIds += fmt.Sprintf("%d", shelfId)
		} else {
			shelvesIds += fmt.Sprintf(",%d", shelfId)
		}
		if mainShelf {
			mainShelfMap[productId] = shelfId
		} else {
			if _, exists := otherShelfMap[productId]; exists {
				otherShelfMap[productId] += fmt.Sprintf("%d", shelfId)
			} else {
				otherShelfMap[productId] += fmt.Sprintf(",%d", shelfId)
			}

		}

	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error in rows: %v", err)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error in rows: %v", err)
	}

	shelvesMap := make(map[int]string)
	query = `SELECT * FROM shelves  WHERE shelves.shelf_id IN (` + shelvesIds + `)`

	rows, err = db.Queryx(query)
	if err != nil {
		return fmt.Errorf("failed to get shelves: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var shelveID int
		var shelfName string
		if err = rows.Scan(&shelveID, &shelfName); err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}
		shelvesMap[shelveID] = shelfName
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error in rows: %v", err)
	}
	idStrs := strings.Split(productIds, ",")
	var ids []int
	for _, idStr := range idStrs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Print("Ошикбка преобзарвания строку в число")
		}
		ids = append(ids, id)
	}
	text := make(map[string]string)
	for _, id := range ids {

		if _, exists := text[shelvesMap[mainShelfMap[id]]]; exists {
			text[shelvesMap[mainShelfMap[id]]] += fmt.Sprintf("%d", shelfId)
		} else {
			otherShelfMap[productId] += fmt.Sprintf(",%d", shelfId)
		}
		fmt.Println(shelvesMap[mainShelfMap[id]] + orderProductMap[id])
	}
	return nil

}
