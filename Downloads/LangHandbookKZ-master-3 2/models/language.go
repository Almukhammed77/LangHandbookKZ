package models

import (
	"time"
)

type Language struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Year        int    `json:"year"`
	Description string `json:"description" gorm:"type:text"`

	// Популярность и статистика
	Popularity  float64 `json:"popularity" gorm:"default:0"`
	Views       int     `json:"views" gorm:"default:0"`
	Rating      float64 `json:"rating" gorm:"default:0"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`

	// Рыночные показатели
	JobCount  int `json:"job_count" gorm:"default:0"`
	SalaryAvg int `json:"salary_avg" gorm:"default:0"`

	// Характеристики языка
	Difficulty string `json:"difficulty" gorm:"default:'Beginner'"`
	Paradigm   string `json:"paradigm"`
	Typing     string `json:"typing"`
	Compiled   bool   `json:"compiled"`

	// Метаданные
	Logo string `json:"logo"`

	// Статистика обучения
	Tutorials int `json:"tutorials" gorm:"default:0"`

	// Тренды
	Trending bool `json:"trending" gorm:"default:false"`

	// Временные метки
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Связи
	Categories []Category `json:"categories,omitempty" gorm:"many2many:language_categories;"`
}

func (l *Language) GetDifficultyColor() string {
	switch l.Difficulty {
	case "Beginner":
		return "#10b981"
	case "Intermediate":
		return "#f59e0b"
	case "Advanced":
		return "#ef4444"
	default:
		return "#6b7280"
	}
}

func (l *Language) GetRatingStars() string {
	stars := ""
	fullStars := int(l.Rating)
	for i := 0; i < fullStars; i++ {
		stars += "★"
	}
	for i := 0; i < 5-fullStars; i++ {
		stars += "☆"
	}
	return stars
}
