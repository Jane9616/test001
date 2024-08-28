package handlers

import (
	"encoding/json"
	"net/http"

	"test0812/models"

	"github.com/gorilla/mux"
)

var users = map[string]models.User{
	"1": {ID: "1", Name: "Jane", Email: "Jane@google.com"},
	"2": {ID: "2", Name: "Shane", Email: "Shane@google.com"},
	"3": {ID: "3", Name: "Jane", Email: "Jane@google.com"},
	"4": {ID: "4", Name: "Shane", Email: "Shane@google.com"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parems := mux.Vars(r)
	id := parems["id"]
	if user, ok := users[id]; ok {
		json.NewEncoder(w).Encode(user)

	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})

	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invail request payload"})
		return
	}
	users[user.ID] = user
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	parems := mux.Vars(r)
	id := parems["id"]
	var updateUser models.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}
	if _, ok := users[id]; ok {
		updateUser.ID = id
		users[id] = updateUser
		json.NewEncoder(w).Encode(updateUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not fount"})
	}
}
