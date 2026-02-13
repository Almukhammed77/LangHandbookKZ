package models

type Language struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Year        int        `json:"year"`
	Description string     `json:"description"`
	Popularity  float64    `json:"popularity"`
	Views       int        `json:"views"`
	Categories  []Category `json:"categories,omitempty" gorm:"many2many:language_categories;"`
}
