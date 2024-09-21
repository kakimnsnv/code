package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	databases "github.com/kakimnsnv/golang-kbtu/assignments/2/databases"
	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"
)

func GetUsersGORM(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	ageStr := r.URL.Query().Get("age")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		http.Error(w, "Invalid page size", http.StatusBadRequest)
		return
	}

	var ageFilter *int
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age filter", http.StatusBadRequest)
			return
		}
		ageFilter = &age
	}

	users, err := databases.GetUsersWithFilterAndPaginationGORM(ageFilter, page, pageSize)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := databases.GetUserByIDGORM(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUserGORM(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := databases.CreateUserGORM(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	json.NewDecoder(r.Body).Decode(&users)
	err := databases.InsertMultipleUsersGORM(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUserGORM(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := databases.UpdateUserGORM(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := databases.DeleteUserGORM(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
