package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var fakeLanguages = []map[string]interface{}{
	{"id": 1, "name": "Go", "year": 2009, "description": "..."},
	{"id": 2, "name": "Python", "year": 1991, "description": "..."},
}

func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(fakeLanguages)

	case http.MethodPost:
		var newLang map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&newLang); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		newID := len(fakeLanguages) + 1
		newLang["id"] = newID
		fakeLanguages = append(fakeLanguages, newLang)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLang)

	default:
		http.Error(w, "Method not allowed (only GET and POST)", http.StatusMethodNotAllowed)
	}
}

func GetLanguageByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed here", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.Trim(r.URL.Path[len("/api/languages/"):], "/")
	if idStr == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(fakeLanguages) {
		http.Error(w, "Not found info", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fakeLanguages[id-1])
}
