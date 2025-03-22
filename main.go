package main

import (
	"log"
	"net/http"
	"day32/api"
)

func main() {
	http.HandleFunc("/api/users", api.UsersHandler)

	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}