package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Routes creates a REST router for the resources resource
func (a api) resourceRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", a.resourceListHandler)    // GET /resources - read a list of resources
	r.Post("/", a.resourceCreateHandler) // POST /resources - create a new resource and peaist it
	r.Put("/", a.resourceDeleteHandler)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(a.resourceCtx) // lets have a resources map, and lets actually load/manipulate
		r.Get("/", a.resourceGetHandler)       // GET /resources/{id} - read a single resource by :id
		r.Put("/", a.resourceUpdateHandler)    // PUT /resources/{id} - update a single resource by :id
		r.Delete("/", a.resourceDeleteHandler) // DELETE /resources/{id} - delete a single resource by :id
		r.Get("/sync", a.resourceSyncHandler)
	})

	return r
}

func (a api) resourceListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resources list of stuff.."))
}

func (a api) resourceCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resources create"))
}

func (a api) resourceGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resource get"))
}

func (a api) resourceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resource update"))
}

func (a api) resourceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resource delete"))
}

func (a api) resourceSyncHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resource sync"))
}
