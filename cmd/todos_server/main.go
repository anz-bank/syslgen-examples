package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	impl "github.com/anz-bank/syslgen-examples/pkg/todos-impl"

	"github.com/anz-bank/syslgen-examples/gen/todos"
)

func main() {
	serviceImpl := impl.NewServiceImpl()
	svcHandler := todos.NewServiceHandler(serviceImpl)
	serviceRouter := todos.NewServiceRouter(svcHandler)
	router := chi.NewRouter()
	serviceRouter.Route(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
