package main

import (
	"log"
	"net/http"

	handlers "github.com/kakimnsnv/golang-kbtu/assignments/2/api/handlers"
	databases "github.com/kakimnsnv/golang-kbtu/assignments/2/databases"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	databases.ConnectSQL()
	databases.ConnectGORM()
	// databases.DropTablesSQL()

	r.HandleFunc("/users", handlers.GetUsersSQL).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserSQL).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUserSQL).Methods("POST")
	r.HandleFunc("/users/bulk", handlers.CreateUsersSQL).Methods("POST")
	r.HandleFunc("/users", handlers.UpdateUserSQL).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserSQL).Methods("DELETE")

	r.HandleFunc("/gorm/users", handlers.GetUsersGORM).Methods("GET")
	r.HandleFunc("/gorm/users/{id}", handlers.GetUserGORM).Methods("GET")
	r.HandleFunc("/gorm/users", handlers.CreateUserGORM).Methods("POST")
	r.HandleFunc("/gorm/users/bulk", handlers.CreateUsersGORM).Methods("POST")
	r.HandleFunc("/gorm/users", handlers.UpdateUserGORM).Methods("PUT")
	r.HandleFunc("/gorm/users/{id}", handlers.DeleteUserGORM).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
