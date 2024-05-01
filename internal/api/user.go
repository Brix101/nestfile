package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Routes creates a REST router for the todos resource
func (a api) resourceRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", a.resourceListHandler)    // GET /todos - read a list of todos
	r.Post("/", a.resourceCreateHandler) // POST /todos - create a new todo and peaist it
	r.Put("/", a.resourceDeleteHandler)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(a.TodoCtx) // lets have a todos map, and lets actually load/manipulate
		r.Get("/", a.resourceGetHandler)       // GET /todos/{id} - read a single todo by :id
		r.Put("/", a.resourceUpdateHandler)    // PUT /todos/{id} - update a single todo by :id
		r.Delete("/", a.resourceDeleteHandler) // DELETE /todos/{id} - delete a single todo by :id
		r.Get("/sync", a.resourceSyncHandler)
	})

	return r
}

func (a api) resourceListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos list of stuff.."))
}

func (a api) resourceCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos create"))
}

func (a api) resourceGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo get"))
}

func (a api) resourceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo update"))
}

func (a api) resourceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo delete"))
}

func (a api) resourceSyncHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo sync"))
}
