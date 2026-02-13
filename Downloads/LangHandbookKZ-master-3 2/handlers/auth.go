package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

// POST /api/register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	// Вариант А — если функция возвращает *models.User, error
	id, err := storage.RegisterUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, `{"error":"username already exists"}`, http.StatusConflict)
		return
	}
	// дальше использовать user
	var user models.User
	storage.DB.First(&user, id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	})
}

// POST /api/login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	token, err := storage.LoginUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
