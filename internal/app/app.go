package app

import (
	"collectorOrder/internal/db"
	"collectorOrder/internal/db/models"
	"collectorOrder/internal/db/query"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("ids")
	IDs := convertStringToInt(idStr)
	orders := getOrdersByID(IDs)
	text := createMessageText(orders)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(text)

}

func convertStringToInt(idStr string) []int {
	idStr = strings.Trim(idStr, ",")

	idSl := strings.Split(idStr, ",")
	var result []int
	for _, id := range idSl {
		num, _ := strconv.Atoi(id)

		result = append(result, num)
	}
	return result
}
func getOrdersByID(IDs []int) map[string][]*models.OrderProductInfo {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()

	orderMap := make(map[string][]*models.OrderProductInfo)
	// Проходимся по данным и добавляем их в карту
	for _, id := range IDs {
		orders, _ := query.GetOrders(sqlDB, id)

		for i, order := range orders {
			if i >= 1 && orders[i].ProductName == orders[i-1].ProductName {
				continue
			}
			orderMap[order.ShelfName] = append(orderMap[order.ShelfName], order)
		}
	}
	return orderMap
}

func createMessageText(orders map[string][]*models.OrderProductInfo) map[string]string {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()

	//Сортируем ключи,так как порядок элементов в map не гарантирован
	keys := make([]string, 0, len(orders))
	for key := range orders {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var textMap = make(map[string]string)

	// Итерируем по отсортированным ключам
	for _, key := range keys {
		for _, ord := range orders[key] {

			textMap[fmt.Sprintf("Стеллаж %s", key)] += fmt.Sprintf("%s (id=%d) заказ %d, %d шт ",
				ord.ProductName, ord.ProductID, ord.OrderID, ord.Quantity)
			shelves, _ := query.GetShelves(sqlDB, ord.ProductID)
			//в случае если продукт находится на нескольких стеллажах
			if len(shelves) > 1 {
				var text string
				for _, shelve := range shelves {
					if shelve.ShelfName != ord.ShelfName {
						text += fmt.Sprintf(" %s", shelve.ShelfName)
					}
				}
				textMap[fmt.Sprintf("Стеллаж %s", key)] += fmt.Sprintf("Доп стелаж%s", text)
			}
		}
	}
	return textMap
}

func getShelvesByID() {

}
