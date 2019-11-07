package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	impl "github.com/anz-bank/syslgen-examples/pkg/todosImpl"

	"github.com/anz-bank/syslgen-examples/gen/todos"
)

const defaultServerAddress = "localhost:8080"

func main() {
	serviceImpl := impl.NewServiceImpl()
	svcHandler := todos.NewServiceHandler(serviceImpl)
	serviceRouter := todos.NewServiceRouter(svcHandler)
	router := chi.NewRouter()
	serviceRouter.Route(router)

	var serverAddress string
	flag.StringVar(&serverAddress, "p", defaultServerAddress, "Specify server address")
	flag.Parse()

	log.Println("Starting Server on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
