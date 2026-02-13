package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		langs := storage.GetAllLanguages("", "")
		json.NewEncoder(w).Encode(langs)
		return
	}

	if r.Method == "POST" {
		var lang models.Language
		if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
		created := storage.CreateLanguage(&lang)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
		return
	}

	http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
}

func LanguageByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/languages/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	if r.Method == "GET" {
		lang := storage.GetLanguageByID(uint(id))
		if lang == nil {
			http.Error(w, `{"error":"language not found"}`, http.StatusNotFound)
			return
		}
		views := lang.Views + 1
		storage.UpdateViews(uint(id), views)
		lang.Views = views
		json.NewEncoder(w).Encode(lang)
		return
	}

	if r.Method == "PUT" {
		var updated models.Language
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
		result := storage.UpdateLanguage(uint(id), &updated)
		if result == nil {
			http.Error(w, `{"error":"language not found"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if r.Method == "DELETE" {
		if storage.DeleteLanguage(uint(id)) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, `{"error":"language not found"}`, http.StatusNotFound)
		}
		return
	}

	http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	results := storage.SearchLanguages(query)
	json.NewEncoder(w).Encode(results)
}

func RatingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var rating models.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	storage.AddRating(&rating)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rating)
}
