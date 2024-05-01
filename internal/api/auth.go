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

func (a api) authRoutes() chi.Router {
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

func (a api) loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var reqBody loginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		a.responseError(w, r, err, 500)
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		a.responseError(w, r, err, 400)
		return
	}

	usr, err := a.userRepo.GetByUsername(ctx, reqBody.Username)
	if err != nil {
		a.responseError(w, r, domain.ErrInvalidCredentials, 401)
		return
	}

	if isValPass := usr.CheckPwd(reqBody.Password); !isValPass {
		a.responseError(w, r, domain.ErrInvalidCredentials, 401)
		return
	}

	token, err := usr.GenerateClaims()
	if err != nil {
		a.logger.Error("failed to generate user claims", zap.Error(err))
		a.responseError(w, r, err, 500)
		return
	}

	data := domain.ResponseUser{
		User: &usr,
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
	a.responseJSON(w, r, data)
}

func (a api) logoutHandler(w http.ResponseWriter, r *http.Request) {

	data := domain.ResponseUser{}
	// Create and set cookies in the response
	cookie := http.Cookie{
		Name:     middlewares.NestfileToken, // Cookie name
		Value:    "",                        // Clear the cookie value for logout
		Path:     "/",                       // Cookie path
		HttpOnly: true,                      // Prevent JavaScript access
		MaxAge:   -1,                        // Set MaxAge to negative value to expire the cookie
	}

	http.SetCookie(w, &cookie)
	a.responseJSON(w, r, data)
}

func (a api) signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sign up stuff"))
}

func (a api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := a.getAuthUser(w, r)

	data := domain.ResponseUser{
		User: user,
	}

	a.responseJSON(w, r, data)
}
