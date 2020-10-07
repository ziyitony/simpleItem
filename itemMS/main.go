package main

import "net/http"

type Item struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	SellerId string  `json:"seller_id"`
}

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Address  string `json:"address"`
}

type ItemDetail struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Seller User    `json:"seller"`
}

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/items", listOrCreateItem)
	http.HandleFunc("/itemdetail", getItemDetail)

	http.ListenAndServe(":55555", nil)
}
