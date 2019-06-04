package main

import (
	"log"
	"net/http"

	impl "github.com/anz-bank/syslgen-examples/todos-impl"

	"github.com/anz-bank/syslgen-examples/todos"
)

func main() {
	serviceImpl := impl.NewServiceImpl()
	svcHandler := todos.NewServiceHandler(serviceImpl)
	serviceRouter := todos.NewServiceRouter(svcHandler)
	router := serviceRouter.Route()

	log.Fatal(http.ListenAndServe(":8080", router))
}
