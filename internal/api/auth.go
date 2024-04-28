package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func (a api) AuthRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	r.Post("/sign-in", a.signInHandler)
	r.Post("/sign-up", a.signUpHandler)
	r.Get("/user", a.getUserHandler)

	return r
}

func (a api) signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign in stuff"))
}

func (a api) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up stuff"))
}

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ResponseUser struct {
	User *User `json:"user"`
}

func (a api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create a user object with null value
	user := &User{
		ID:        1,
		Username:  "admin",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data := ResponseUser{
		User: user,
	}

	// Convert data object to JSON
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.Write(jsonResponse)
}
