package main

import "net/http"

const maxLength = 1000

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/items", listOrCreateItem)
	http.HandleFunc("/itemdetail", getItemDetail)

	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/users", listOrCreateUser)
	http.HandleFunc("/userid/", getUserById)

	http.ListenAndServe(":12345", nil)
}
