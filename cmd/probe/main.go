package main

import (
	"log"
	"net/http"

	"github.com/zealllot/proctor-fixtures/api"
)

func main() {
	http.HandleFunc("/api/users", api.ListUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
