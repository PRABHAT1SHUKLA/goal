package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Checked bool   `json:"checked"`
}

var todos []Todo
var idCounter = 1

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTodo.ID = idCounter
	idCounter++
	todos = append(todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	idStr := keys.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Deleted"))
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			createTodo(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
