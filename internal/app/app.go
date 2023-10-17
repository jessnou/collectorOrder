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
			shelves, _ := query.GetShelves(sqlDB, ord.ProductID)
			//в случае если продукт находится на нескольких стеллажах
			if len(shelves) == 1 {
				text += "\n \n"
			} else {
				var shelvesText string
				for _, shelve := range shelves {
					if shelve.ShelfName != ord.ShelfName {
						shelvesText += fmt.Sprintf(" %s", shelve.ShelfName)
					}
				}
				text += fmt.Sprintf("\nДоп стелаж%s \n\n", shelvesText)

			}
		}
	}
	return text
}
