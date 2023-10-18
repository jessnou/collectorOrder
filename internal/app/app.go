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
	"strings"
)

func ParseCommandLineArgs() (string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Нет переданных аргументов ")
	}

	result := convertToString(args)
	return result, nil
}

func GetOrdersShelvesProducts(ids string) ([]models.Shelves, []models.ProductShelf, []models.OrderProduct) {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()
	var productIdsInt []int
	var shelvesIdsInt []int

	orders, _ := query.GetOrders(sqlDB, ids)
	for _, ord := range orders {
		productIdsInt = append(productIdsInt, ord.ProductID)
	}

	productIds := convertIntToString(productIdsInt)
	products, _ := query.GetProductShelf(sqlDB, productIds)
	for _, pr := range products {
		shelvesIdsInt = append(shelvesIdsInt, pr.ShelfId)
	}

	shelvesIds := convertIntToString(shelvesIdsInt)
	shelves, _ := query.GetShelves(sqlDB, shelvesIds)
	return shelves, products, orders
}

func CreateMap(shelves []*models.Shelves, products []*models.ProductShelf,
	orders []*models.OrderProduct) map[string][]*models.OrderProductInfo {
	orderMap := make(map[string][]*models.OrderProductInfo)
	otherShelves := make(map[int][]*string)
	for _, s := range shelves {
		for _, pr := range products {
			if pr.ShelfId == s.ShelveID && !pr.MainShelf {
				otherShelves[pr.ProductId] = append(otherShelves[pr.ProductId], &s.ShelfName)
			}
		}
	}

	for _, ord := range orders {

		for _, s := range shelves {

			for _, pr := range products {
				if pr.ShelfId == s.ShelveID && ord.ProductID == pr.ProductId && pr.MainShelf {
					orderMap[s.ShelfName] = append(orderMap[s.ShelfName], &models.OrderProductInfo{
						OrderID:     ord.OrderID,
						ProductName: pr.ProductName,
						ProductID:   pr.ProductId,
						Quantity:    ord.Quantity,
						OtherShelf:  otherShelves[pr.ProductId],
					})

				}
			}
		}
	}

	return orderMap
}

func CreateMessageCmd(orders map[string][]*models.OrderProductInfo) {
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
			//в случае если продукт находится на нескольких стеллажах
			if ord.OtherShelf == nil {
				text += "\n \n"
			} else {
				var shelvesText string
				for _, shelve := range ord.OtherShelf {

					shelvesText += fmt.Sprintf(" %s", *shelve)

				}
				text += fmt.Sprintf("\nДоп стеллаж%s \n\n", shelvesText)

			}

		}
	}
	fmt.Println(text)
}

func convertToString(args []string) string {
	result := fmt.Sprintf("(%s)", strings.Join(args, ","))
	return result
}
func convertIntToString(numbers []int) string {
	stringNumbers := make([]string, len(numbers))
	for i, num := range numbers {
		stringNumbers[i] = fmt.Sprint(num)
	}
	result := convertToString(stringNumbers)
	return result
}
