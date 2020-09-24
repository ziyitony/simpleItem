package main

import "net/http"

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/item", createItem)
	http.HandleFunc("/items", listItems)

	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/userid/", getUserById)

	http.ListenAndServe(":12345", nil)
}
