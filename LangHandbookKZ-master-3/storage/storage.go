package storage

import (
	"strings"

	"github.com/Almukhammed77/LangHandbookKZ/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("langhandbook.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	DB = db
	DB.AutoMigrate(&models.Language{}, &models.Category{}, &models.Rating{})
}

func CreateLanguage(lang *models.Language) *models.Language {
	DB.Create(lang)
	return lang
}

func GetAllLanguages() []*models.Language {
	var langs []*models.Language
	DB.Preload("Categories").Find(&langs)
	return langs
}

func GetLanguageByID(id uint) *models.Language {
	var lang models.Language
	if DB.Preload("Categories").First(&lang, id).Error != nil {
		return nil
	}
	return &lang
}

func UpdateLanguage(id uint, updated *models.Language) *models.Language {
	lang := GetLanguageByID(id)
	if lang == nil {
		return nil
	}
	DB.Model(lang).Updates(updated)
	return lang
}

func DeleteLanguage(id uint) bool {
	return DB.Delete(&models.Language{}, id).Error == nil
}

func SearchLanguages(query string) []*models.Language {
	var langs []*models.Language
	tx := DB.Preload("Categories")
	if query != "" {
		q := "%" + strings.ToLower(query) + "%"
		tx = tx.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", q, q).
			Or("id IN (SELECT language_id FROM language_categories lc JOIN categories c ON lc.category_id = c.id WHERE LOWER(c.name) LIKE ?)", q)
	}
	tx.Find(&langs)
	return langs
}

func AddRating(rating *models.Rating) {
	DB.Create(rating)
	var avg float64
	DB.Model(&models.Rating{}).Where("language_id = ?", rating.LanguageID).Select("AVG(score)").Scan(&avg)
	DB.Model(&models.Language{}).Where("id = ?", rating.LanguageID).Update("popularity", avg)
}

func UpdateViews(langID uint, views int) {
	DB.Model(&models.Language{}).Where("id = ?", langID).Update("views", views)
}
