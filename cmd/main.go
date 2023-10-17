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
	m := app.GetOrdersByID(numbers)

	text := app.CreateMessageCmd(m)
	fmt.Println(text)

}
