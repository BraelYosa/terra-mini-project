package model

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	User_id     uint   `gorm:"primaryKey"`
	UserEmail   string `gorm:"column:user_mail" json:"user_mail"`
	UserSurname string `gorm:"column:user_name" json:"user_name"`
	UserPass    string `gorm:"column:user_pass" json:"user_pass"` //hash password
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type FormData struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type LoginRequest struct {
	UserMail string `json:"username"` // Match the frontend key ("username")
	Password string `json:"password"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
