package main

import "net/http"

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/item", createItem)
	http.HandleFunc("/items", listItems)

	http.ListenAndServe(":55555", nil)
}
