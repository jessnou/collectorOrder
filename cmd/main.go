package main

import (
	"collectorOrder/internal/app"
)

func main() {
	//numbers, err := app.ParseCommandLineArgs()
	//if err != nil {
	//	log.Fatal(err)
	//}
	ids := "10,11,14,15"
	_, _, _, _, _, _ = app.GetOrdersShelvesProducts(ids)

	//app.CreateMessageCmd(m)

}
