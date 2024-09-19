package restapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	dbRepo "github.com/kakimnsnv/golang-kbtu/assignments/2/databases"
	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"
)

// MARK: SQL
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.Connect()
	if err != nil {
		log.Println("Error occured with db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Println("Error while quering from db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Age)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.Connect()
	if err != nil {
		log.Println("Error occured with db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error while decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		log.Println("Error while inserting into db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.Connect()
	if err != nil {
		log.Println("Error occured with db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error while decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
	if err != nil {
		log.Println("Error while updating db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.Connect()
	if err != nil {
		log.Println("Error occured with db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM users WHERE id=$1", strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		log.Println("Error while deleting from db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User deleted"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.Connect()
	if err != nil {
		log.Println("Error occured with db", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE id=$1", strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		log.Println("Error while quering from db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Age)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// MARK: GORM
func GetUsersGorm(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.ConnectGorm()
	if err != nil {
		log.Printf("Error while connecting to db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func CreateUserGorm(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.ConnectGorm()
	if err != nil {
		log.Printf("Error while connecting to db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error while decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserGorm(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.ConnectGorm()
	if err != nil {
		log.Printf("Error while connecting to db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error while decoding request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserGorm(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.ConnectGorm()
	if err != nil {
		log.Printf("Error while connecting to db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	db.Delete(&user, strings.Split(r.URL.Path, "/")[3])
	w.Write([]byte("User deleted"))
}

func GetUserGorm(w http.ResponseWriter, r *http.Request) {
	db, err := dbRepo.ConnectGorm()
	if err != nil {
		log.Printf("Error while connecting to db")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	db.First(&user, strings.Split(r.URL.Path, "/")[3])
	json.NewEncoder(w).Encode(user)
}
