package models

type Rating struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	LanguageID uint `json:"language_id"`
	Score      int  `json:"score"` // 1–5
}
