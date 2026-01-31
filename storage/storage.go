package storage

import (
	"sync"

	"github.com/Almukhammed77/LangHandbookKZ/models"
)

import "strings"

var (
	languages = make(map[uint]*models.Language)
	mu        sync.RWMutex
	nextID    uint = 1
)

func CreateLanguage(lang *models.Language) *models.Language {
	mu.Lock()
	defer mu.Unlock()
	lang.ID = nextID
	nextID++
	languages[lang.ID] = lang
	return lang
}

func GetAllLanguages() []*models.Language {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]*models.Language, 0, len(languages))
	for _, lang := range languages {
		result = append(result, lang)
	}
	return result
}

func GetLanguageByID(id uint) *models.Language {
	mu.RLock()
	defer mu.RUnlock()
	return languages[id]
}

func UpdateLanguage(id uint, updated *models.Language) *models.Language {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := languages[id]; !exists {
		return nil
	}
	updated.ID = id
	languages[id] = updated
	return updated
}

func DeleteLanguage(id uint) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := languages[id]; exists {
		delete(languages, id)
		return true
	}
	return false
}

func SearchLanguages(query string) []*models.Language {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]*models.Language, 0)
	for _, lang := range languages {
		if query == "" ||
			containsIgnoreCase(lang.Name, query) ||
			containsIgnoreCase(lang.Description, query) {
			result = append(result, lang)
		}
	}
	return result
}

func containsIgnoreCase(s, substr string) bool {
	// простая реализация без strings.ToLower для скорости
	return len(substr) == 0 || strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}