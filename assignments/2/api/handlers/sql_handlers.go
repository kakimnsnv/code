package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	databases "github.com/kakimnsnv/golang-kbtu/assignments/2/databases"
	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"

	"github.com/gorilla/mux"
)

func GetUsersSQL(w http.ResponseWriter, r *http.Request) {
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

	users, err := databases.GetUsersWithFilterAndPaginationSQL(ageFilter, page, pageSize)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := databases.GetUserByIDSQL(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUserSQL(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := databases.CreateUserSQL(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUsersSQL(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	json.NewDecoder(r.Body).Decode(&users)
	err := databases.InsertMultipleUsersSQL(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func UpdateUserSQL(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := databases.UpdateUserSQL(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := databases.DeleteUserSQL(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
