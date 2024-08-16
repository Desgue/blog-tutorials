package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Desgue/blog-tutorials/todo-api/handler"
	"github.com/Desgue/blog-tutorials/todo-api/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost user=develop password=develop dbname=todos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todoRepo := repository.NewTodoRepository(db)
	todoHandler := handler.NewTodoHandler(todoRepo)

	r := mux.NewRouter()
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
