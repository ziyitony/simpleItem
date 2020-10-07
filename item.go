package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ziyitony/simpleItem/domain"
)

var (
	itemMutex sync.Mutex
	itemDB    []*domain.Item
)

func init() {
	item1 := &domain.Item{
		Id:       "m001",
		Name:     "iphone",
		Price:    50000,
		SellerId: "u001",
	}

	item2 := &domain.Item{
		Id:       "m002",
		Name:     "t-shirt",
		Price:    2000,
		SellerId: "u002",
	}

	item3 := &domain.Item{
		Id:       "m003",
		Name:     "sofa",
		Price:    30000,
		SellerId: "u003",
	}

	itemDB = []*domain.Item{item1, item2, item3}
}

func helloItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hello":"item"}`))
}

func listOrCreateItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var item domain.Item
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

// return domain.ItemDetail as the response
func getItemDetail(w http.ResponseWriter, r *http.Request) {
	itemMutex.Lock()
	defer itemMutex.Unlock()

	itemdetails := make([]*domain.ItemDetail, len(itemDB))
	// loop item DB
	for i, item := range itemDB {
		user := getUserByIdLocally(item.SellerId)
		if user == nil {
			// return error
			http.Error(w, "no such seller", http.StatusInternalServerError)
			return
		}
		detail := &domain.ItemDetail{
			Id:    item.Id,
			Name:  item.Name,
			Price: item.Price,
			// get the user info by userID
			Seller: *user,
		}
		itemdetails[i] = detail
	}

	// unmarshal and response
	data, err := json.Marshal(itemdetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
