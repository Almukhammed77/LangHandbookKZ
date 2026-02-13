package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/Almukhammed77/LangHandbookKZ/storage"
	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "super-secret-key-2026"

func WebRegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	if username == "" || password == "" {
		tmpl.Execute(w, map[string]string{"Error": "Заполните поля"})
		return
	}

	_, err := storage.RegisterUser(username, password)
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Логин занят"})
		return
	}

	// Автовход после регистрации
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(jwtSecret))

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenStr,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func WebLoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	user, err := storage.LoginUser(username, password)
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Неверный логин/пароль"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(jwtSecret))

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenStr,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
