package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Brix101/nestfile/internal/domain"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (a api) UserRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", a.userListHandler)    // GET /users - read a list of users
	r.Post("/", a.userCreateHandler) // POST /users - create a new user and persist it
	r.Put("/", a.userDeleteHandler)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(a.TodoCtx) // lets have a users map, and lets actually load/manipulate
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

type createUserRequest struct {
	Username string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required,min=6"` // Minimum length: 6
}

func (a api) userCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var reqBody createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUsr := domain.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
	}

	usr, err := a.userRepo.Create(ctx, &newUsr)
	if err != nil {
		a.logger.Error("failed to create user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resJSON, err := json.Marshal(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
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
