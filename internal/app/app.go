package app

//import (
//	"collectorOrder/internal/db"
//	"collectorOrder/internal/db/models"
//	"collectorOrder/internal/db/query"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func ParseCommandLineArgs() (string, []int, error) {
//	args := os.Args[1:]
//	if len(args) == 0 {
//		log.Fatal("Нет переданных аргументов ")
//	}
//	delimiter := ","
//	argStr := strings.Join(args, " ")
//	argList := strings.Split(argStr, delimiter)
//	var numbers []int
//	for _, arg := range argList {
//		num, err := strconv.Atoi(arg)
//		if err != nil {
//			return "", nil, fmt.Errorf("Ошибка преобразования аргумента %s в число: %v\n", arg, err)
//		}
//		numbers = append(numbers, num)
//	}
//	result := convertToString(args)
//	return result, numbers, nil
//}
//
//func GetOrdersShelvesProducts(ids string) (map[int][]*models.OrderProduct, map[int]*models.Product,
//	map[int][]*models.ProductShelf, map[int]*models.Shelves, []int, []int) {
//	sqlDB, _ := db.GetDBConn()
//	defer sqlDB.Close()
//	var productIdsInt []int
//	var shelvesIdsInt []int
//
//	orderProducts, _ := query.GetOrderProducts(sqlDB, ids)
//	for prID := range orderProducts {
//		productIdsInt = append(productIdsInt, prID)
//
//	}
//
//	productIds := convertIntToString(productIdsInt)
//
//	products, _ := query.GetProduct(sqlDB, productIds)
//
//	productShelves, _ := query.GetProductShelf(sqlDB, productIds)
//	for psID := range productShelves {
//		shelvesIdsInt = append(shelvesIdsInt, psID)
//	}
//
//	shelvesIds := convertIntToString(shelvesIdsInt)
//	shelves, _ := query.GetShelves(sqlDB, shelvesIds)
//	return orderProducts, products, productShelves, shelves, productIdsInt, shelvesIdsInt
//}
//
//func CreateMessageCmd(orderProducts map[int][]*models.OrderProduct,
//	products map[int]*models.Product,
//	productShelves map[int][]*models.ProductShelf,
//	shelves map[int]*models.Shelves, ids []int, productIds []int, shelvesIds []int) {
//	for _, ordPr := range orderProducts {
//		for _, op := range ordPr {
//			fmt.Printf("Стеллаж %s", shelves[1].ShelfName)
//		}
//
//	}
//}
//
//func convertToString(args []string) string {
//	result := fmt.Sprintf("%s", strings.Join(args, ","))
//	return result
//}
//func convertIntToString(numbers []int) string {
//	stringNumbers := make([]string, len(numbers))
//	for i, num := range numbers {
//		stringNumbers[i] = fmt.Sprint(num)
//	}
//	result := convertToString(stringNumbers)
//	return result
//}
