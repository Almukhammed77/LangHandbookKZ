package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Color       string `json:"color"`
	SortOrder   int    `json:"sort_order" gorm:"default:0"`

	Languages []Language `json:"languages,omitempty" gorm:"many2many:language_categories;"`
}

func (c *Category) GetIcon() string {
	if c.Icon != "" {
		return c.Icon
	}
	switch c.Name {
	case "Web":
		return "🌐"
	case "Mobile":
		return "📱"
	case "Desktop":
		return "💻"
	case "System":
		return "⚙️"
	case "DataScience":
		return "📊"
	case "MachineLearning":
		return "🤖"
	case "GameDev":
		return "🎮"
	case "Enterprise":
		return "🏢"
	default:
		return "📁"
	}
}
