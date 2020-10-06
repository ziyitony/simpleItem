package main

import "net/http"

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/items", listOrCreateItem)
	http.HandleFunc("/itemdetail", getItemDetail)

	http.ListenAndServe(":55555", nil)
}
