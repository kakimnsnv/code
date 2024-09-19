package databases

import (
	"log"

	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() (*gorm.DB, error) {
	connstr := "user=kakimbekn dbname=golang-kbtu password=Sadasa@2015 host=localhost sslmode=disable"
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	log.Println("Auto migration completed")
}

func CreateUserGorm(db *gorm.DB, name string, age int) {
	user := models.User{Name: name, Age: age}
	db.Create(&user)
	log.Println("User created successfully: ", user)
}

func GetUsersGorm(db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	for _, user := range users {
		log.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
