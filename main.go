package main

import "net/http"

const maxLength = 1000

func main() {
	http.HandleFunc("/helloitem", helloItem)
	http.HandleFunc("/items", listOrCreateItem)
	http.HandleFunc("/itemdetail", listItemDetail)

	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/userid/", getUserById)
	print("test")

	http.ListenAndServe(":12345", nil)
	print("this is a test")
}
