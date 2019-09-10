package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	database "./db"
	handleFunc "./handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	defer database.Close()
	r := mux.NewRouter()

	// Todo endpoints

	r.HandleFunc("/api/todos", handleFunc.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos", handleFunc.GetAllTodos).Methods("GET")
	r.HandleFunc("/api/todos/{todoId}", handleFunc.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos/{todoId}", handleFunc.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/api/todos/{todoId}", handleFunc.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":81", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"DELETE", "GET", "PATCH", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "application/json"}),
	)(r)))
}
