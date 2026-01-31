package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Almukhammed77/LangHandbookKZ/concurrency"
	"github.com/Almukhammed77/LangHandbookKZ/handlers"
	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

func main() {
	storage.CreateLanguage(&models.Language{
		Name:        "Go",
		Year:        2009,
		Description: "Статически типизированный компилируемый язык от Google",
		Popularity:  9.2,
		Categories:  []models.Category{{Name: "Системное"}, {Name: "Concurrency"}},
	})

	storage.CreateLanguage(&models.Language{
		Name:        "Python",
		Year:        1991,
		Description: "Интерпретируемый язык общего назначения",
		Popularity:  9.8,
		Categories:  []models.Category{{Name: "Скрипты"}, {Name: "Data Science"}},
	})

	// Запускаем фоновый счётчик просмотров (goroutine + канал)
	concurrency.StartViewCounter()
	log.Println("Фоновый счётчик просмотров запущен")

	// Роуты
	http.HandleFunc("/api/languages", handlers.LanguagesHandler)
	http.HandleFunc("/api/languages/", LanguageByIDWithViewsHandler) // ← наш обработчик с views
	http.HandleFunc("/api/search", handlers.SearchHandler)

	// Простой welcome для проверки
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "LangHandbookKZ API работает")
		fmt.Fprintln(w, "Доступные эндпоинты:")
		fmt.Fprintln(w, "  GET  /api/languages")
		fmt.Fprintln(w, "  POST /api/languages")
		fmt.Fprintln(w, "  GET  /api/languages/{id}   ← показывает views")
		fmt.Fprintln(w, "  PUT  /api/languages/{id}")
		fmt.Fprintln(w, "  DELETE /api/languages/{id}")
		fmt.Fprintln(w, "  GET  /api/search?q=query")
	})

	fmt.Println("Сервер запущен → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Обработчик для GET /api/languages/{id} с подсчётом просмотров
func LanguageByIDWithViewsHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем часть пути после /api/languages/
	path := strings.TrimPrefix(r.URL.Path, "/api/languages/")
	if path == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	// Парсим id как uint
	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	// Получаем язык (функция возвращает только *models.Language или nil)
	lang := storage.GetLanguageByID(uint(id))

	if lang == nil {
		http.Error(w, "Язык не найден", http.StatusNotFound)
		return
	}

	// Увеличиваем счётчик просмотров асинхронно
	concurrency.AddView(lang.ID)

	// Формируем ответ с полем views
	response := struct {
		*models.Language
		Views int `json:"views"`
	}{
		Language: lang,
		Views:    concurrency.GetViewsCount(lang.ID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
