package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/Brix101/nestfile/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func (a api) AuthRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	r.Post("/login", a.loginHandler)
	r.Post("/logout", a.logoutHandler)
	r.Post("/sign-up", a.signUpHandler)
	r.Get("/user", a.getUserHandler)

	return r
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"` // Minimum length: 6
}

type ResponseUser struct {
	User *domain.User `json:"user"`
}

func (a api) loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var reqBody loginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	usr, err := a.userRepo.GetByUsername(ctx, reqBody.Username)
	if err != nil {
		http.Error(w, domain.ErrInvalidCredentials.Error(), 401)
		return
	}

	if isValPass := usr.CheckPwd(reqBody.Password); !isValPass {
		http.Error(w, domain.ErrInvalidCredentials.Error(), 401)
		return
	}

	token, err := usr.GenerateClaims()
	if err != nil {
		a.logger.Error("failed to generate user claims", zap.Error(err))
		http.Error(w, err.Error(), 500)
		return
	}

	data := ResponseUser{
		User: &usr,
	}

	resJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Create and set cookies in the response
	cookie := http.Cookie{
		Name:     middlewares.NestfileToken, // Cookie name
		Value:    token,                     // Cookie value (you can customize this)
		Path:     "/",                       // Cookie path
		HttpOnly: true,                      // Prevent JavaScript access
		// You can set more attributes like Expires, MaxAge, Secure, etc. as needed.
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)

}

func (a api) logoutHandler(w http.ResponseWriter, r *http.Request) {
	data := ResponseUser{
		User: nil,
	}

	resJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Create and set cookies in the response
	cookie := http.Cookie{
		Name:     middlewares.NestfileToken, // Cookie name
		Value:    "",                        // Clear the cookie value for logout
		Path:     "/",                       // Cookie path
		HttpOnly: true,                      // Prevent JavaScript access
		MaxAge:   -1,                        // Set MaxAge to negative value to expire the cookie
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (a api) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up stuff"))
}

func (a api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	// TODO move this section to its own  getter
	var user *domain.User
	usrCtx := ctx.Value(middlewares.UserCtxKey{})
	if usrCtx != nil {
		item, ok := usrCtx.(*domain.UserClaims)
		if ok {
			usr, err := a.userRepo.GetByID(ctx, int64(item.Sub))
			if err == nil {
				user = &usr
			}
		}
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
	w.Write(jsonResponse)
}
