package models

type Language struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Year        int        `json:"year"`
	Description string     `json:"description"`
	Popularity  float64    `json:"popularity"`
	Categories  []Category `json:"categories,omitempty"`
}