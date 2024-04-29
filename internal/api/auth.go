package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/go-chi/chi/v5"
)

func (a api) AuthRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	r.Post("/login", a.loginHandler)
	r.Post("/sign-up", a.signUpHandler)
	r.Get("/user", a.getUserHandler)

	return r
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"` // Minimum length: 6
}

func (a api) loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var reqBody loginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usr, err := a.userRepo.GetByUsername(ctx, reqBody.Username)
	if err != nil {
		http.Error(w, domain.ErrInvalidCredentials.Error(), 401)
		return
	}

	log.Println(usr)
	if validatePass := usr.CheckPassword(reqBody.Password); !validatePass {
		http.Error(w, domain.ErrInvalidCredentials.Error(), 401)
		return
	}

	resJSON, err := json.Marshal(usr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)

}

func (a api) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up stuff"))
}

type ResponseUser struct {
	User *domain.User `json:"user"`
}

func (a api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create a user object with null value
	// user := &User{
	// 	ID:        1,
	// 	Username:  "admin",
	// 	Password:  "password",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }
	data := ResponseUser{
		User: nil,
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
