package main

import "net/http"

func main() {
	http.HandleFunc("/hellouser", helloUser)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/userid/", getUserById)

	http.ListenAndServe(":44444", nil)
}
