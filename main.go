package main

import (
	"log"
	"net/http"

	"github.com/Almukhammed77/LangHandbookKZ/handlers"
)

func main() {
	http.HandleFunc("/api/languages", handlers.LanguagesHandler)
	http.HandleFunc("/api/languages/", handlers.GetLanguageByID)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
