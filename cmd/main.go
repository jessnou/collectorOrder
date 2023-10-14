package main

import "collectorOrder/internal/app"

func main() {
	id, _ := app.ParseCommandLineArgs()
	orders := app.Query(id)
	app.CreateMessageCmd(orders)
}
