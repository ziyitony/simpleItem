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
