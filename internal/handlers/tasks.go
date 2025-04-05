package handlers

import (
	"encoding/json"
	"habit-tracker-api/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	r.HandleFunc("/lists/{listId}/tasks", getTasks).Methods("GET")
	r.HandleFunc("/lists/{listId}/tasks", createTask).Methods("POST")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var listTasks []models.Task
	for _, task := range models.Tasks {
		if task.ListID == params["listId"] {
			listTasks = append(listTasks, task)
		}
	}
	json.NewEncoder(w).Encode(listTasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	json.NewDecoder(r.Body).Decode(&newTask)
	models.Tasks = append(models.Tasks, newTask)
	w.WriteHeader(http.StatusCreated)
}
