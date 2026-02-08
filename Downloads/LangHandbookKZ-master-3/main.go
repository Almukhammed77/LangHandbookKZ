package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Almukhammed77/LangHandbookKZ/concurrency"
	"github.com/Almukhammed77/LangHandbookKZ/handlers"
	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

func main() {
	storage.InitDB()
	log.Println("The database has been initialized. (langhandbook.db)")

	concurrency.StartViewCounter()
	log.Println("The background view counter has been started.")

	var count int64
	storage.DB.Model(&models.Language{}).Count(&count)
	if count == 0 {
		log.Println("Adding seed languages...")

		storage.CreateLanguage(&models.Language{
			Name:        "Go",
			Year:        2009,
			Description: "Статически типизированный компилируемый язык от Google",
			Popularity:  0,
			Categories:  []models.Category{{Name: "Системное"}, {Name: "Concurrency"}},
		})

		storage.CreateLanguage(&models.Language{
			Name:        "Python",
			Year:        1991,
			Description: "Интерпретируемый язык общего назначения",
			Popularity:  0,
			Categories:  []models.Category{{Name: "Скрипты"}, {Name: "Data Science"}},
		})

		log.Println("Initial data added")
	}

	http.HandleFunc("/api/languages", handlers.LanguagesHandler)
	http.HandleFunc("/api/languages/", handlers.LanguageByIDHandler)
	http.HandleFunc("/api/search", handlers.SearchHandler)
	http.HandleFunc("/api/ratings", handlers.RatingsHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		langs := storage.GetAllLanguages()

		data := struct {
			Count     int
			Languages []*models.Language
		}{
			Count:     len(langs),
			Languages: langs,
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html.tmpl"))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	cors := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	fmt.Println("Server running on → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", cors(http.DefaultServeMux)))
}
