package main

import (
	"collectorOrder/internal/app"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	http.HandleFunc("/get", app.GetOrders)

	_ = http.ListenAndServe(":8080", nil)

}
