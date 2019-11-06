package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	impl "github.com/anz-bank/syslgen-examples/pkg/todos-impl"

	"github.com/anz-bank/syslgen-examples/gen/todos"
)

const serverPort = ":8080"

func main() {
	serviceImpl := impl.NewServiceImpl()
	svcHandler := todos.NewServiceHandler(serviceImpl)
	serviceRouter := todos.NewServiceRouter(svcHandler)
	router := chi.NewRouter()
	serviceRouter.Route(router)

	log.Println("Starting Server on localhost" + serverPort)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
