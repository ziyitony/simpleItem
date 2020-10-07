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
	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/users", listOrCreateUser)
	http.HandleFunc("/userid/", getUserById)

	http.ListenAndServe(":44444", nil)
}
