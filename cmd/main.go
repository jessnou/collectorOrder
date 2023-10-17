package main

import (
	"collectorOrder/internal/app"
	"fmt"
)

func main() {
	//numbers, err := app.ParseCommandLineArgs()
	//if err != nil {
	//	log.Fatal(err)
	//}
	var ids = []int{10, 11, 14, 15}

	orders := app.GetOrdersByID(ids)

	input := app.CreateMessageCmd(orders)

	fmt.Println(input)

}
