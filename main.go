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
		Description: "Google's statically typed compiled language",
		Popularity:  9.2,
		Categories:  []models.Category{{Name: "Системное"}, {Name: "Concurrency"}},
	})

	storage.CreateLanguage(&models.Language{
		Name:        "Python",
		Year:        1991,
		Description: "General-purpose interpreted language",
		Popularity:  9.8,
		Categories:  []models.Category{{Name: "Scripts"}, {Name: "Data Science"}},
	})

	concurrency.StartViewCounter()
	log.Println("The background view counter has been started.")

	// Оборачиваем handlers в CORS
	http.Handle("/api/languages", cors(http.HandlerFunc(handlers.LanguagesHandler)))
	http.Handle("/api/languages/", cors(http.HandlerFunc(LanguageByIDWithViewsHandler)))
	http.Handle("/api/search", cors(http.HandlerFunc(handlers.SearchHandler)))

	// Фронтенд (статические файлы)
	http.Handle("/", cors(http.FileServer(http.Dir("./static"))))

	fmt.Println("server run on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CORS middleware — простой и правильный
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Для preflight-запросов (OPTIONS) сразу отвечаем OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LanguageByIDWithViewsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/languages/")
	if path == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	lang := storage.GetLanguageByID(uint(id))
	if lang == nil {
		http.Error(w, "Язык не найден", http.StatusNotFound)
		return
	}

	concurrency.AddView(lang.ID)

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
