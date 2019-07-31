package todos

//
//    THIS IS AUTOGENERATED BY syslgen
//

import (
	"github.com/go-chi/chi"
)

// Router interface for Todos
type Router interface {
	Route(router *chi.Mux)
}

// ServiceRouter for Todos API
type ServiceRouter struct {
	svcHandler *ServiceHandler
}

// NewServiceRouter for Todos
func NewServiceRouter(svcHandler *ServiceHandler) *ServiceRouter {
	return &ServiceRouter{svcHandler}
}

// Route ...
func (r *ServiceRouter) Route(router *chi.Mux) {
	router.Get("/comments", r.svcHandler.GetCommentsHandler)
	router.Get("/posts", r.svcHandler.GetPostsHandler)
	router.Get("/todos/{id}", r.svcHandler.GetTodosIDHandler)
	router.Post("/comments", r.svcHandler.PostCommentsHandler)
}
