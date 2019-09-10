package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	database "./db"
	_handlers "./handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	defer database.Close()
	r := mux.NewRouter()

	// Todo endpoints

	r.HandleFunc("/api/todos", _handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos", _handlers.GetAllTodos).Methods("GET")
	r.HandleFunc("/api/todos/{todoId}", _handlers.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos/{todoId}", _handlers.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/api/todos/{todoId}", _handlers.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":81", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"DELETE", "GET", "PATCH", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "application/json"}),
	)(r)))
}
