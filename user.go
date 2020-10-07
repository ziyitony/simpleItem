package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ziyitony/simpleItem/domain"
)

var (
	userMutex sync.Mutex
	userDB    []*domain.User
)

func init() {
	user1 := &domain.User{
		Id:       "u001",
		Nickname: "nico",
		Address:  "yokohama",
	}

	user2 := &domain.User{
		Id:       "u002",
		Nickname: "tony",
		Address:  "tokyo",
	}

	user3 := &domain.User{
		Id:       "u003",
		Nickname: "gogo",
		Address:  "chugoku",
	}

	userDB = []*domain.User{user1, user2, user3}
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hello":"user"}`))
}

func listOrCreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user domain.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userMutex.Lock()
		defer userMutex.Unlock()

		if len(userDB) >= maxLength-1 {
			http.Error(w, "database is full", http.StatusInternalServerError)
			return
		}
		user.Id = fmt.Sprintf("u%03d", len(userDB)+1)
		userDB = append(userDB, &user)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":"succeed"}`))
	case http.MethodGet:
		userMutex.Lock()
		defer userMutex.Unlock()

		data, err := json.Marshal(userDB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		http.Error(w, "unsupported HTTP method", http.StatusMethodNotAllowed)
	}
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

	http.Error(w, "no such user", http.StatusBadRequest)
}

func getUserByIdLocally(userID string) *domain.User {
	userMutex.Lock()
	defer userMutex.Unlock()

	for _, user := range userDB {
		if user.Id == userID {
			return user
		}
	}
	return nil
}
