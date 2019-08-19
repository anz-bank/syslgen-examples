package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"

	impl "github.com/anz-bank/syslgen-examples/todos-impl"

	"github.com/anz-bank/syslgen-examples/todos"
)

func main() {
	router := chi.NewRouter()
	todos.RouteServices(router,impl.NewServiceImpl())

	log.Fatal(http.ListenAndServe(":8080", router))
}
