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

	// Миграция
	err = DB.AutoMigrate(&models.Language{}, &models.Category{}, &models.User{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	// Создаем категории если их нет
	createCategories()

	// Обновляем языки
	updateLanguages()
}

func createCategories() {
	var count int64
	DB.Model(&models.Category{}).Count(&count)

	if count == 0 {
		categories := []models.Category{
			{Name: "Web", Description: "Веб-разработка", Icon: "🌐", Color: "#3b82f6", SortOrder: 1},
			{Name: "Mobile", Description: "Мобильная разработка", Icon: "📱", Color: "#10b981", SortOrder: 2},
			{Name: "Desktop", Description: "Десктопные приложения", Icon: "💻", Color: "#6366f1", SortOrder: 3},
			{Name: "System", Description: "Системное программирование", Icon: "⚙️", Color: "#8b5cf6", SortOrder: 4},
			{Name: "DataScience", Description: "Наука о данных", Icon: "📊", Color: "#ec4899", SortOrder: 5},
			{Name: "MachineLearning", Description: "Машинное обучение", Icon: "🤖", Color: "#f43f5e", SortOrder: 6},
			{Name: "GameDev", Description: "Разработка игр", Icon: "🎮", Color: "#f97316", SortOrder: 7},
			{Name: "Enterprise", Description: "Корпоративная разработка", Icon: "🏢", Color: "#64748b", SortOrder: 8},
		}

		for _, cat := range categories {
			DB.Create(&cat)
		}
	}
}

func updateLanguages() {
	var languages []models.Language
	DB.Find(&languages)

	for _, lang := range languages {
		updates := make(map[string]interface{})

		switch lang.Name {
		case "JavaScript":
			updates["difficulty"] = "Beginner"
			updates["rating"] = 4.7
			updates["review_count"] = 1250
			updates["job_count"] = 25000
			updates["salary_avg"] = 85000
			updates["trending"] = true
			updates["tutorials"] = 8

			var categories []models.Category
			DB.Where("name IN ?", []string{"Web", "Desktop"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "Python":
			updates["difficulty"] = "Beginner"
			updates["rating"] = 4.9
			updates["review_count"] = 2100
			updates["job_count"] = 22000
			updates["salary_avg"] = 90000
			updates["trending"] = true
			updates["tutorials"] = 7

			var categories []models.Category
			DB.Where("name IN ?", []string{"DataScience", "MachineLearning", "Web"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "Java":
			updates["difficulty"] = "Intermediate"
			updates["rating"] = 4.5
			updates["review_count"] = 980
			updates["job_count"] = 20000
			updates["salary_avg"] = 95000
			updates["trending"] = false
			updates["tutorials"] = 8

			var categories []models.Category
			DB.Where("name IN ?", []string{"Enterprise", "Mobile"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "Go":
			updates["difficulty"] = "Intermediate"
			updates["rating"] = 4.8
			updates["review_count"] = 650
			updates["job_count"] = 8000
			updates["salary_avg"] = 110000
			updates["trending"] = true
			updates["tutorials"] = 9

			var categories []models.Category
			DB.Where("name IN ?", []string{"System", "DevOps"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "TypeScript":
			updates["difficulty"] = "Intermediate"
			updates["rating"] = 4.8
			updates["review_count"] = 720
			updates["job_count"] = 15000
			updates["salary_avg"] = 95000
			updates["trending"] = true
			updates["tutorials"] = 9

			var categories []models.Category
			DB.Where("name IN ?", []string{"Web", "Desktop"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "C#":
			updates["difficulty"] = "Intermediate"
			updates["rating"] = 4.6
			updates["review_count"] = 580
			updates["job_count"] = 12000
			updates["salary_avg"] = 90000
			updates["trending"] = false
			updates["tutorials"] = 7

			var categories []models.Category
			DB.Where("name IN ?", []string{"GameDev", "Desktop", "Enterprise"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "Rust":
			updates["difficulty"] = "Advanced"
			updates["rating"] = 4.9
			updates["review_count"] = 420
			updates["job_count"] = 3000
			updates["salary_avg"] = 120000
			updates["trending"] = true
			updates["tutorials"] = 7

			var categories []models.Category
			DB.Where("name IN ?", []string{"System", "Embedded"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "C++":
			updates["difficulty"] = "Advanced"
			updates["rating"] = 4.5
			updates["review_count"] = 890
			updates["job_count"] = 10000
			updates["salary_avg"] = 100000
			updates["trending"] = false
			updates["tutorials"] = 7

			var categories []models.Category
			DB.Where("name IN ?", []string{"System", "GameDev", "Embedded"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)

		case "Swift":
			updates["difficulty"] = "Intermediate"
			updates["rating"] = 4.7
			updates["review_count"] = 510
			updates["job_count"] = 5000
			updates["salary_avg"] = 105000
			updates["trending"] = true
			updates["tutorials"] = 7

			var categories []models.Category
			DB.Where("name IN ?", []string{"Mobile", "Desktop"}).Find(&categories)
			DB.Model(&lang).Association("Categories").Replace(categories)
		}

		if len(updates) > 0 {
			DB.Model(&lang).Updates(updates)
		}
	}
}

// ОСНОВНЫЕ ФУНКЦИИ ДЛЯ ФИЛЬТРАЦИИ

func GetAllLanguages(category string, sort string) []*models.Language {
	var langs []*models.Language
	tx := DB.Preload("Categories")

	if category != "" && category != "all" {
		tx = tx.Joins("JOIN language_categories lc ON lc.language_id = languages.id").
			Joins("JOIN categories c ON c.id = lc.category_id").
			Where("c.name = ?", category)
	}

	if sort != "" {
		switch sort {
		case "popularity":
			tx = tx.Order("views DESC")
		case "rating":
			tx = tx.Order("rating DESC")
		case "name":
			tx = tx.Order("name ASC")
		case "name_desc":
			tx = tx.Order("name DESC")
		case "salary":
			tx = tx.Order("salary_avg DESC")
		case "newest":
			tx = tx.Order("year DESC")
		case "oldest":
			tx = tx.Order("year ASC")
		default:
			tx = tx.Order("views DESC")
		}
	}

	tx.Find(&langs)
	return langs
}

func FilterLanguages(filters map[string]interface{}) []*models.Language {
	var languages []*models.Language
	query := DB.Preload("Categories")

	if category, ok := filters["category"]; ok && category != "" && category != "all" {
		query = query.Joins("JOIN language_categories lc ON lc.language_id = languages.id").
			Joins("JOIN categories c ON c.id = lc.category_id").
			Where("c.name = ?", category)
	}

	if difficulty, ok := filters["difficulty"]; ok && difficulty != "" && difficulty != "all" {
		query = query.Where("difficulty = ?", difficulty)
	}

	if search, ok := filters["search"]; ok && search != "" {
		q := "%" + strings.ToLower(search.(string)) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", q, q)
	}

	if minRating, ok := filters["min_rating"]; ok {
		query = query.Where("rating >= ?", minRating)
	}

	query.Find(&languages)
	return languages
}

func GetLanguageByID(id uint) *models.Language {
	var lang models.Language
	DB.Preload("Categories").First(&lang, id)
	return &lang
}

func GetAllCategories() []*models.Category {
	var categories []*models.Category
	DB.Order("sort_order").Find(&categories)
	return categories
}

func SearchLanguages(query string) []*models.Language {
	var langs []*models.Language
	tx := DB.Preload("Categories")
	if query != "" {
		q := "%" + strings.ToLower(query) + "%"
		tx = tx.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", q, q)
	}
	tx.Find(&langs)
	return langs
}

func UpdateViews(langID uint, views int) {
	DB.Model(&models.Language{}).Where("id = ?", langID).Update("views", views)
}

// ФУНКЦИИ ДЛЯ ПОЛЬЗОВАТЕЛЕЙ

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
