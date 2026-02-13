package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"-"`
	Email      string    `json:"email" gorm:"default:''"`
	FullName   string    `json:"full_name" gorm:"default:''"`
	Bio        string    `json:"bio" gorm:"type:text;default:''"`
	Location   string    `json:"location" gorm:"default:''"`
	Avatar     string    `json:"avatar" gorm:"default:'/static/avatars/default.png'"`
	Role       string    `json:"role" gorm:"default:'user'"`
	Level      int       `json:"level" gorm:"default:1"`
	Experience int       `json:"experience" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
}
