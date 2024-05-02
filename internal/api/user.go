package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func (a api) userRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", a.userListHandler)    // GET /users - read a list of users
	r.Post("/", a.userCreateHandler) // POST /users - create a new user and persist it
	r.Put("/", a.userDeleteHandler)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(a.userCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", a.userGetHandler)       // GET /users/{id} - read a single user by :id
		r.Put("/", a.userUpdateHandler)    // PUT /users/{id} - update a single user by :id
		r.Delete("/", a.userDeleteHandler) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

func (a api) userListHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	data, err := a.userRepo.GetAll(ctx)
	if err != nil {
		a.logger.Error("failed to fetch all user from database", zap.Error(err))
		a.responseError(w, r, err, 00)
		return
	}

	a.responseJSON(w, r, data)
}

type createUserDTO struct {
	Username string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required,min=6"` // Minimum length: 6
}

func (a api) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var reqDTO createUserDTO
	if err := json.NewDecoder(r.Body).Decode(&reqDTO); err != nil {
		a.responseError(w, r, err, 500)
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqDTO); err != nil {
		a.responseError(w, r, err, 400)
		return
	}

	usr := domain.User{
		Username: reqDTO.Username,
		Password: reqDTO.Password,
	}

	if err := usr.HashPwd(); err != nil {
		a.logger.Error("failed to hash user password", zap.Error(err))
		a.responseError(w, r, err, 500)
		return
	}

	usr, err := a.userRepo.Create(ctx, &usr)
	if err != nil {
		a.logger.Error("failed to create user", zap.Error(err))
		a.responseError(w, r, err, 500)
		return
	}

	a.responseJSON(w, r, usr)
}

func (a api) userGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user get"))
}

func (a api) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user update"))
}

func (a api) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user delete"))
}
