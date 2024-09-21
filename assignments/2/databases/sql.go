package databases

import (
	"database/sql"
	"log"

	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"

	_ "github.com/lib/pq"
)

var dbSQL *sql.DB

func ConnectSQL() *sql.DB {
	connStr := "user=kakimbekn dbname=golang-kbtu password=Sadasa@2015 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * 60 * 1e9)

	dbSQL = db
	createTableSQL()
	return dbSQL
}

func createTableSQL() {
	userTable := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        age INT NOT NULL
    );`

	profileTable := `CREATE TABLE IF NOT EXISTS profiles (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        bio TEXT,
        profile_picture_url TEXT
    );`

	dbSQL.Exec(userTable)
	dbSQL.Exec(profileTable)
}

func DropTablesSQL() {
	dbSQL.Exec("DROP TABLE users CASCADE")
	dbSQL.Exec("DROP TABLE profiles CASCADE")
}

func CreateUserSQL(user *models.User) error {
	err := dbSQL.QueryRow(
		"INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id",
		user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		return err
	}

	_, err = dbSQL.Exec("INSERT INTO profiles (user_id, bio, profile_picture_url) VALUES ($1, $2, $3)",
		user.ID, user.Profile.Bio, user.Profile.ProfilePictureURL)
	return err
}

func GetUserByIDSQL(id uint) (*models.User, error) {
	var user models.User
	err := dbSQL.QueryRow("SELECT id, name, age FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}

	err = dbSQL.QueryRow("SELECT bio, profile_picture_url FROM profiles WHERE user_id=$1", user.ID).
		Scan(&user.Profile.Bio, &user.Profile.ProfilePictureURL)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserSQL(user *models.User) error {
	_, err := dbSQL.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}

	_, err = dbSQL.Exec("UPDATE profiles SET bio=$1, profile_picture_url=$2 WHERE user_id=$3",
		user.Profile.Bio, user.Profile.ProfilePictureURL, user.ID)
	return err
}

func DeleteUserSQL(id uint) error {
	_, err := dbSQL.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func GetAllUsersSQL() ([]models.User, error) {
	rows, err := dbSQL.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		err = dbSQL.QueryRow("SELECT bio, profile_picture_url FROM profiles WHERE user_id=$1", user.ID).
			Scan(&user.Profile.Bio, &user.Profile.ProfilePictureURL)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func InsertMultipleUsersSQL(users []models.User) error {
	tx, err := dbSQL.Begin()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := tx.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", user.Name, user.Age).
			Scan(&user.ID)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec("INSERT INTO profiles (user_id, bio, profile_picture_url) VALUES ($1, $2, $3)",
			user.ID, user.Profile.Bio, user.Profile.ProfilePictureURL)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func GetUsersWithFilterAndPaginationSQL(ageFilter *int, page, pageSize int) ([]models.User, error) {
	var rows *sql.Rows
	var err error
	offset := (page - 1) * pageSize

	if ageFilter != nil {
		rows, err = dbSQL.Query("SELECT id, name, age FROM users WHERE age = $1 OFFSET $2 LIMIT $3", *ageFilter, offset, pageSize)
	} else {
		rows, err = dbSQL.Query("SELECT id, name, age FROM users OFFSET $1 LIMIT $2", offset, pageSize)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		err = dbSQL.QueryRow("SELECT bio, profile_picture_url FROM profiles WHERE user_id=$1", user.ID).
			Scan(&user.Profile.Bio, &user.Profile.ProfilePictureURL)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
