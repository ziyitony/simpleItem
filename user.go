package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Address  string `json:"address"`
}

var (
	userMutex sync.Mutex
	userDB    []*User
)

func init() {
	user1 := &User{
		Id:       "u001",
		Nickname: "nico",
		Address:  "yokohama",
	}

	user2 := &User{
		Id:       "u002",
		Nickname: "tony",
		Address:  "tokyo",
	}

	user3 := &User{
		Id:       "u003",
		Nickname: "gogo",
		Address:  "chugoku",
	}

	userDB = []*User{user1, user2, user3}
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hello":"user"}`))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	if user.Id == "" {
		user.Id = fmt.Sprintf("u%03d", len(userDB)+1)
	}
	userDB = append(userDB, &user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result":"succeed"}`))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method is allowed", http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	data, err := json.Marshal(userDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method is allowed", http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	userID := r.URL.Path[8:]
	for _, user := range userDB {
		if user.Id == userID {
			json.NewEncoder(w).Encode(*user)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
}
