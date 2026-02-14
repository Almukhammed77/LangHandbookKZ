package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Password   string    `json:"-"` // "-" скрывает поле в JSON
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	Bio        string    `json:"bio" gorm:"type:text"`
	Location   string    `json:"location"`
	Role       string    `json:"role" gorm:"default:'user'"`
	Level      int       `json:"level" gorm:"default:1"`
	Experience int       `json:"experience" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
