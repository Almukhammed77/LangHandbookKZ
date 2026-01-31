package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Almukhammed77/LangHandbookKZ/handlers"
	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

func main() {
	// Тестовые данные
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

	// Роуты
	http.HandleFunc("/api/languages", handlers.LanguagesHandler)
	http.HandleFunc("/api/languages/", handlers.LanguageByIDHandler) // слеш обязателен
	http.HandleFunc("/api/search", handlers.SearchHandler)

	// Простой welcome для проверки
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "LangHandbookKZ API работает")
		fmt.Fprintln(w, "Доступные эндпоинты:")
		fmt.Fprintln(w, "  GET  /api/languages")
		fmt.Fprintln(w, "  POST /api/languages")
		fmt.Fprintln(w, "  GET  /api/languages/{id}")
		fmt.Fprintln(w, "  PUT  /api/languages/{id}")
		fmt.Fprintln(w, "  DELETE /api/languages/{id}")
		fmt.Fprintln(w, "  GET  /api/search?q=query")
	})

	fmt.Println("Сервер запущен → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}