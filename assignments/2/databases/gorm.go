package databases

import (
	"log"

	models "github.com/kakimnsnv/golang-kbtu/assignments/2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbGORM *gorm.DB

func ConnectGORM() *gorm.DB {
	dsn := "user=kakimbekn dbname=golang-kbtu password=Sadasa@2015 host=localhost sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Profile{})
	dbGORM = db
	return dbGORM
}

func CreateUserGORM(user *models.User) error {
	return dbGORM.Create(user).Error
}

func GetUserByIDGORM(id uint) (*models.User, error) {
	var user models.User
	if err := dbGORM.Preload("Profile").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserGORM(user *models.User) error {
	return dbGORM.Save(user).Error
}

func DeleteUserGORM(id uint) error {
	return dbGORM.Delete(&models.User{}, id).Error
}

func GetAllUsersGORM() ([]models.User, error) {
	var users []models.User
	if err := dbGORM.Preload("Profile").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func InsertMultipleUsersGORM(users []models.User) error {
	tx := dbGORM.Begin()
	for _, user := range users {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func GetUsersWithFilterAndPaginationGORM(ageFilter *int, page, pageSize int) ([]models.User, error) {
	var users []models.User
	query := dbGORM.Preload("Profile")

	if ageFilter != nil {
		query = query.Where("age = ?", *ageFilter)
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
