package databases

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connstr := "user=kakimbekn dbname=golang-kbtu password=Sadasa@2015 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users( id SERIAL PRIMARY KEY, name VARCHAR (50) NOT NULL, age INT NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return
	}
	log.Println("Table created successfully")
}

func InsertUser(db *sql.DB, name string, age int) {
	query := `INSERT INTO users(name, age) VALUES($1, $2)`
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return
	}
	log.Println("User inserted successfully")
}

func GetUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return
		}
		log.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
