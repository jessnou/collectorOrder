package main

import (
	"collectorOrder/internal/app"
)

func main() {
	//numbers, err := app.ParseCommandLineArgs()
	//if err != nil {
	//	log.Fatal(err)
	//}
	ids := "(10,11,14,15)"
	o, s, p := app.GetOrdersShelvesProducts(ids)
	m := app.CreateMap(o, s, p)
	app.CreateMessageCmd(m)

}
