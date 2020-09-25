package main

import "net/http"

func main() {
	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/users", listOrCreateUser)
	http.HandleFunc("/userid/", getUserById)

	http.ListenAndServe(":44444", nil)
}
