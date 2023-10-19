package main

import (
	"collectorOrder/internal/db"
	"collectorOrder/internal/db/query"
	_ "github.com/lib/pq"
)

func main() {
	//numbers, err := app.ParseCommandLineArgs()
	//if err != nil {
	//	log.Fatal(err)
	//}
	sqlDB, _ := db.GetDBConn()
	defer sqlDB.Close()
	ids := "10,11,14,15"
	_ = query.GetOrderProducts(sqlDB, ids)
	//app.CreateMessageCmd(m)

}
