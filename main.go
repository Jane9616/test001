package main

import (
	"fmt"
	"net/http"

	"test0812/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	fmt.Println("server is running on port 8800")
	// http.ListenAndServe(":8080", r)
	if err := http.ListenAndServe(":8800", r); err != nil {
		fmt.Println("Error running server", err)
	}
}
