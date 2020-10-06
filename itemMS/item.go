package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const maxLength = 1000

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

var (
	itemMutex sync.Mutex
	itemDB    []*Item
)

func init() {
	item1 := &Item{
		Id:       "m001",
		Name:     "iphone",
		Price:    50000,
		SellerId: "u001",
	}

	item2 := &Item{
		Id:       "m002",
		Name:     "t-shirt",
		Price:    2000,
		SellerId: "u002",
	}

	item3 := &Item{
		Id:       "m003",
		Name:     "sofa",
		Price:    30000,
		SellerId: "u003",
	}

	itemDB = []*Item{item1, item2, item3}
}

func helloItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hello":"item"}`))
}

func listOrCreateItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		itemMutex.Lock()
		defer itemMutex.Unlock()

		if len(itemDB) >= maxLength-1 {
			http.Error(w, "database is full", http.StatusInternalServerError)
			return
		}
		item.Id = fmt.Sprintf("m%03d", len(itemDB)+1)
		itemDB = append(itemDB, &item)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":"succeed"}`))
	case http.MethodGet:
		itemMutex.Lock()
		defer itemMutex.Unlock()

		data, err := json.Marshal(itemDB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		http.Error(w, "unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func listItemDetail(w http.ResponseWriter, r *http.Request) {
	itemMutex.Lock()
	defer itemMutex.Unlock()

	itemDetails := make([]*ItemDetail, len(itemDB))
	for i, item := range itemDB {
		// HTTP get /userid/{id}
		resp, err := http.Get("http://localhost:44444/userid/" + item.SellerId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		}
		defer resp.Body.Close()

		fmt.Println(resp.Body)
		var user User
		err = json.NewDecoder(resp.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		}

		itemDetails[i] = &ItemDetail{
			Id:     item.Id,
			Name:   item.Name,
			Price:  item.Price,
			Seller: user,
		}
	}

	data, err := json.Marshal(itemDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
