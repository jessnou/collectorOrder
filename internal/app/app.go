package app

import (
	"collectorOrder/internal/db"
	"collectorOrder/internal/db/query"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseCommandLineArgs() (string, error) {
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
			return "", fmt.Errorf("Ошибка преобразования аргумента %s в число: %v\n", arg, err)
		}
		numbers = append(numbers, num)
	}
	result := convertToString(args)
	return result, nil
}

func TextMap(ids string) map[string]string {
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()

	orderProductMap, productIds := query.GetOrderProducts(sqlDB, ids)
	productMap := query.GetProductMap(sqlDB, productIds)
	mainShelfMap, otherShelfMap, shelvesIds := query.GetProductShelvesMaps(sqlDB, productIds)
	shelvesMap := query.GetShelvesMap(sqlDB, shelvesIds)

	otherShelvesName := make(map[int]string)
	for key := range otherShelfMap {
		idsStrM := strings.Split(otherShelfMap[key], ",")
		var idsM []int
		for _, idStrM := range idsStrM {
			idM, err := strconv.Atoi(idStrM)
			if err != nil {
				fmt.Print("Ошибка преобзарвания строку в число")
			}
			idsM = append(idsM, idM)
		}
		for _, id := range idsM {

			otherShelvesName[key] += " " + shelvesMap[id]
		}

	}
	idStrProd := strings.Split(productIds, ",")
	idStrOrder := strings.Split(ids, ",")
	var idsPr []int
	var idsOrd []int
	for _, idStrO := range idStrOrder {
		id, err := strconv.Atoi(idStrO)
		if err != nil {
			fmt.Print("Ошибка преобзарвания строку в число")
		}
		idsOrd = append(idsOrd, id)
	}

	for _, idStrP := range idStrProd {
		id, err := strconv.Atoi(idStrP)
		if err != nil {
			fmt.Print("Ошибка преобзарвания строку в число")
		}
		idsPr = append(idsPr, id)
	}
	text := make(map[string]string)

	for _, idPr := range idsPr {

		for _, idOrd := range idsOrd {
			if _, exists := orderProductMap[fmt.Sprintf("%d,%d", idOrd, idPr)]; exists {
				if _, exists = text[shelvesMap[mainShelfMap[idPr]]]; !exists {
					text[shelvesMap[mainShelfMap[idPr]]] = ""
				}

				if strings.Contains(text[shelvesMap[mainShelfMap[idPr]]], fmt.Sprintf("%s %s", productMap[idPr], orderProductMap[fmt.Sprintf("%d,%d", idOrd, idPr)])) {
					continue
				}

				text[shelvesMap[mainShelfMap[idPr]]] += fmt.Sprintf("%s %s", productMap[idPr], orderProductMap[fmt.Sprintf("%d,%d", idOrd, idPr)])
				if _, ok := otherShelfMap[idPr]; ok {
					text[shelvesMap[mainShelfMap[idPr]]] += fmt.Sprintf("\nДоп стеллаж %s\n\n", otherShelvesName[idPr])

				} else {
					text[shelvesMap[mainShelfMap[idPr]]] += "\n\n"
				}
			}
		}

	}
	return text
}
func CreateMessageCmd(text map[string]string) {

	keys := make([]string, 0, len(text))
	for key := range text {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Стеллаж %s\n%s", key, text[key])
	}

}

func convertToString(args []string) string {
	result := fmt.Sprintf("%s", strings.Join(args, ","))
	return result
}
