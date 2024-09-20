package models

type User struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Profile Profile `json:"profile" gorm:"foreignKey:UserID"`
}
