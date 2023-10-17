package main

import (
	"collectorOrder/internal/app"
	"fmt"
	"log"
)

func main() {
	numbers, err := app.ParseCommandLineArgs()
	if err != nil {
		log.Fatal(err)
	}

	orders := app.GetOrdersByID(numbers)

	input := app.CreateMessageCmd(orders)

	fmt.Println(input)

}
