package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Almukhammed77/LangHandbookKZ/concurrency"
	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		langs := storage.GetAllLanguages()
		json.NewEncoder(w).Encode(langs)
	case http.MethodPost:
		var lang models.Language
		if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
			http.Error(w, `{"error":"invalid JSON"}`, http.StatusBadRequest)
			return
		}
		created := storage.CreateLanguage(&lang)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func LanguageByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/languages/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, `{"error":"invalid ID"}`, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		lang := storage.GetLanguageByID(uint(id))
		if lang == nil {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
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

		json.NewEncoder(w).Encode(response)

	case http.MethodPut:
		var updated models.Language
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, `{"error":"invalid JSON"}`, http.StatusBadRequest)
			return
		}
		result := storage.UpdateLanguage(uint(id), &updated)
		if result == nil {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(result)
	case http.MethodDelete:
		if storage.DeleteLanguage(uint(id)) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		}
	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	results := storage.SearchLanguages(query)
	json.NewEncoder(w).Encode(results)
}
