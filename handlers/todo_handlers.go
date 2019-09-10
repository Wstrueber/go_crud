package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "../db"
	models "../models"
	"github.com/gorilla/mux"
)

// CreateTodo creates todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var todo models.Todo
	err := decoder.Decode(&todo)
	if err != nil {
		return
	}
	fmt.Println(todo)
	rows, err := database.DB.Query(`insert into todo
	(content, status, order_number)
	select $1, $2, $3
	returning id`, todo.Content, false, todo.OrderNumber)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var count, id int
	for rows.Next() {
		count++
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if count == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.ID = id
	todo.Status = false
	json.NewEncoder(w).Encode(todo)

}

// GetAllTodos gets todo by {todoId}
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todos []models.Todo
	var id int
	var content string
	var status bool
	var orderNumber int
	rows, err := database.DB.Query("SELECT * FROM todo ORDER BY order_number ASC")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &content, &status, &orderNumber)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, models.Todo{ID: id, Content: content, Status: status, OrderNumber: orderNumber})
	}

	json.NewEncoder(w).Encode(todos)
}

// GetTodo gets todo by {todoId}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	todoID, err := strconv.Atoi(params["todoId"])
	if err != nil {
		return
	}
	var todo models.Todo
	rows, err := database.DB.Query(`SELECT * from todo WHERE
	id=$1`, todoID)
	defer rows.Close()
	var count, id int
	var content string
	var status bool
	var orderNumber int

	for rows.Next() {
		count++
		err = rows.Scan(&id, &content, &status, &orderNumber)
		if err != nil {
			log.Fatal(err)
		}
	}
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	todo = models.Todo{ID: id, Content: content, Status: status, OrderNumber: orderNumber}

	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo deletes todo by {todoId}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := strconv.Atoi(params["todoId"])
	_, err = database.DB.Query(`DELETE FROM todo WHERE id=$1`, todoID)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateTodo updates content, status, orderNumber by {todoId}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var todo models.Todo
	todoID, _ := strconv.Atoi(params["todoId"])
	todo.ID = todoID
	err := decoder.Decode(&todo)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(todo.OrderNumber)
	rows, err := database.DB.Query(`UPDATE todo
	SET content=$1,
	status=$2,
	order_number=$3
	WHERE id=$4
	returning
	id, content, status, order_number`, todo.Content, todo.Status, todo.OrderNumber, todoID)
	if err != nil {
		return
	}
	var id int
	var content string
	var status bool
	var orderNumber int

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &content, &status, &orderNumber)

		if err != nil {
			log.Fatal(err)
		}
	}

	todo = models.Todo{ID: id, Content: content, Status: status, OrderNumber: orderNumber}

	json.NewEncoder(w).Encode(todo)
}
