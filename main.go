package main

import (
	"net/http"
)

type Item struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price float64	`json:"price"`
	SellerId string `json:"seller_id"`
}

func main() {
	http.HandleFunc("/hello", helloItem)
	//http.HandleFunc("/item", createItem)
	http.ListenAndServe(":12345", nil)
}

func helloItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hello":"item"}`))
}

//func createItem(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte`{"":""}`)
//	}
//}