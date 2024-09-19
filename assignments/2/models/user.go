package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
	Age  int
}
