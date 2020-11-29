package main

import (
	"homework/week02/getting"
	"homework/week02/storage"
	"homework/week02/users"
	"net/http"
)

func main() {
	var userStorage users.Repository
	userStorage = storage.NewMockUserStorage()
	service := getting.NewService(userStorage)

	// HTTP server
	http.Handle("/user", http.HandlerFunc(getting.MakeGetUserEndpoint(service)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start HTTP server")
	}
}
