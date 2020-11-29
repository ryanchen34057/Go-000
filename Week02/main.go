package main

import (
	"Go-000/Week02/getting"
	"Go-000/Week02/storage"
	"Go-000/Week02/users"
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
