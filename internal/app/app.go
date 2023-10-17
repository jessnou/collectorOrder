package app

import (
	"collectorOrder/internal/db"
	"collectorOrder/internal/db/models"
	"collectorOrder/internal/db/query"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseCommandLineArgs() ([]int, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Нет переданных аргументов ")
	}
	delimiter := ","
	argStr := strings.Join(args, " ")
	argList := strings.Split(argStr, delimiter)
	var numbers []int
	for _, arg := range argList {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования аргумента %s в число: %v\n", arg, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func GetOrdersByID(IDs []int) map[string][]*models.OrderProductInfo {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()
	orderMap := make(map[string][]*models.OrderProductInfo)

	for _, id := range IDs {
		orders, _ := query.GetOrders(sqlDB, id)
		for _, ord := range orders {

			products, _ := query.GetProductShelf(sqlDB, ord.ProductID)
			var otherShelves []string
			for _, pr := range products {

				if !pr.MainShelf {
					shelf, _ := query.GetShelf(sqlDB, pr.ShelfId)

					otherShelves = append(otherShelves, shelf.ShelfName)

				}
			}
			for _, pr := range products {

				if pr.MainShelf {
					shelf, _ := query.GetShelf(sqlDB, pr.ShelfId)
					orderMap[shelf.ShelfName] = append(orderMap[shelf.ShelfName], &models.OrderProductInfo{
						OrderID:     ord.OrderID,
						ProductName: pr.ProductName,
						ProductID:   pr.ProductId,
						Quantity:    ord.Quantity,
						OtherShelf:  otherShelves,
					})
				}
			}
		}
	}

	return orderMap
}
func CreateMessageCmd(orders map[string][]*models.OrderProductInfo) string {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()

	//Сортируем ключи,так как порядок элементов в map не гарантирован
	keys := make([]string, 0, len(orders))
	for key := range orders {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var text string
	// Итерируем по отсортированным ключам
	for _, key := range keys {
		text += fmt.Sprintf("===Стеллаж %s \n", key)
		for _, ord := range orders[key] {
			text += fmt.Sprintf("%s (id=%d)\nзаказ %d, %d шт",
				ord.ProductName, ord.ProductID, ord.OrderID, ord.Quantity)
			////в случае если продукт находится на нескольких стеллажах
			if ord.OtherShelf == nil {
				text += "\n \n"
			} else {
				var shelvesText string
				for _, shelve := range ord.OtherShelf {

					shelvesText += fmt.Sprintf(" %s", shelve)

				}
				text += fmt.Sprintf("\nДоп стеллаж%s \n\n", shelvesText)

			}
		}
	}
	return text
}
