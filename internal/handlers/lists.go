package handlers

import (
	"encoding/json"
	"habit-tracker-api/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterListRoutes(r *mux.Router) {
	r.HandleFunc("/lists", getLists).Methods("GET")
	r.HandleFunc("/lists", createList).Methods("POST")
	r.HandleFunc("/lists/{id}", deleteList).Methods("DELETE")
}

func getLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Lists)
}

func createList(w http.ResponseWriter, r *http.Request) {
	var newList models.List
	json.NewDecoder(r.Body).Decode(&newList)
	models.Lists = append(models.Lists, newList)
	w.WriteHeader(http.StatusCreated)
}

func deleteList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, list := range models.Lists {
		if list.ID == params["id"] {
			models.Lists = append(models.Lists[:i], models.Lists[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
