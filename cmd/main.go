package main

import (
	"collectorOrder/internal/app"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	ids, err := app.ParseCommandLineArgs()
	if err != nil {
		log.Fatal(err)
	}

	mapText := app.TextMap(ids)
	app.CreateMessageCmd(mapText)

}
