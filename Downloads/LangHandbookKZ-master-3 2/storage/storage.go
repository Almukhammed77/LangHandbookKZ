package storage

import (
	"strings"
	"time"

	"github.com/Almukhammed77/LangHandbookKZ/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("langhandbook.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	DB = db
	DB.AutoMigrate(&models.Language{}, &models.Category{}, &models.Rating{}, &models.User{})
}

func CreateLanguage(lang *models.Language) *models.Language {
	DB.Create(lang)
	return lang
}

func GetAllLanguages(category string, sort string) []*models.Language {
	var langs []*models.Language
	tx := DB.Preload("Categories")

	if category != "" {
		q := "%" + strings.ToLower(category) + "%"
		tx = tx.Joins("JOIN language_categories lc ON lc.language_id = languages.id").
			Joins("JOIN categories c ON c.id = lc.category_id").
			Where("LOWER(c.name) LIKE ?", q)
	}

	if sort != "" {
		tx = tx.Order(sort)
	}

	tx.Find(&langs)
	return langs
}

func GetLanguageByID(id uint) *models.Language {
	var lang models.Language
	DB.Preload("Categories").First(&lang, id)
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

func RegisterUser(username, password string) (*models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:   username,
		Password:   string(hashed),
		Email:      username + "@example.com",
		FullName:   username,
		Role:       "user",
		Level:      1,
		Experience: 0,
		CreatedAt:  time.Now(),
	}

	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(username, password string) (*models.User, error) {
	var user models.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) *models.User {
	var user models.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func UpdateUserProfile(userID uint, fullName, email, bio, location string) error {
	updates := map[string]interface{}{
		"full_name": fullName,
		"email":     email,
		"bio":       bio,
		"location":  location,
	}
	return DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}

func UpdatePassword(userID uint, newPassword string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return DB.Model(&models.User{}).Where("id = ?", userID).Update("password", string(hashed)).Error
}
