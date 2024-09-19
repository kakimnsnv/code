package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/kakimnsnv/golang-kbtu/assignments/2/restapi"
)

func main() {
	// MARK: POSTGRES SQL
	// db, err := dbRepo.Connect()
	// if err != nil {
	// 	log.Fatalf("Error connecting to database: %v", err)
	// }

	// dbRepo.CreateTable(db)
	// dbRepo.InsertUser(db, "John Doe", 25)
	// dbRepo.InsertUser(db, "Jane Doe", 30)
	// dbRepo.GetUsers(db)

	// MARK: GORM PORSTGRES
	// db, err := dbRepo.ConnectGorm()
	// if err != nil {
	// 	log.Fatalf("Error connecting to database: %v", err)
	// }

	// dbRepo.AutoMigrate(db)
	// dbRepo.CreateUserGorm(db, "John Doe", 25)
	// dbRepo.CreateUserGorm(db, "Jane Doe", 30)
	// dbRepo.GetUsersGorm(db)

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	// MARK: gorm handlers
	r.HandleFunc("/users/gorm", handlers.GetUsersGorm).Methods("GET")
	r.HandleFunc("/users/gorm", handlers.CreateUserGorm).Methods("POST")
	r.HandleFunc("/users/gorm/{id}", handlers.GetUserGorm).Methods("GET")
	r.HandleFunc("/users/gorm/{id}", handlers.UpdateUserGorm).Methods("PUT")
	r.HandleFunc("/users/gorm/{id}", handlers.DeleteUserGorm).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
