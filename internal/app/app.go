package app

import (
	"collectorOrder/internal/db"
	"collectorOrder/internal/db/models"
	"collectorOrder/internal/db/query"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
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

func Query(IDs []int) [][]*models.OrderProductInfo {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()
	var ordersS [][]*models.OrderProductInfo
	for _, id := range IDs {
		orders, _ := query.Get(sqlDB, id)
		ordersS = append(ordersS, orders)
	}

	return ordersS
}

func CreateMessageCmd(orders [][]*models.OrderProductInfo) {
	for _, order := range orders {
		for _, ord := range order {
			fmt.Println(ord)
		}

	}
}
