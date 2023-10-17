package main

import (
	"collectorOrder/internal/app"
	"log"
)

func main() {
	numbers, err := app.ParseCommandLineArgs()
	if err != nil {
		log.Fatal(err)
	}
	o, s, p := app.GetOrdersShelvesProducts(numbers)
	m := app.CreateMap(o, s, p)
	app.CreateMessageCmd(m)

}
